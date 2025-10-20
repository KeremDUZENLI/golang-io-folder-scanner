package scanner

import (
	"os"
	"path/filepath"
)

func ListFolders(root string) []string {
	folders := make([]string, 0, 64)
	stack := []string{root}

	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		folders = append(folders, dir)

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		folderNames := make([]string, 0, len(entries))
		for _, entry := range entries {
			if entry.IsDir() {
				folderNames = append(folderNames, entry.Name())
			}
		}

		sortStrings(folderNames)

		for i := len(folderNames) - 1; i >= 0; i-- {
			stack = append(stack, filepath.Join(dir, folderNames[i]))
		}
	}

	return folders
}
