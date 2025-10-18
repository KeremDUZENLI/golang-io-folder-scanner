package scanner

import (
	"os"
	"path/filepath"
)

func ListFolders(path string) ([]string, error) {
	folders := make([]string, 0, 64)
	stack := []string{path}

	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		folders = append(folders, dir)

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() {
				stack = append(stack, filepath.Join(dir, e.Name()))
			}
		}
	}

	return folders, nil
}

func ListFiles(folders []string) ([]string, error) {
	files := []string{}
	for _, dir := range folders {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !e.IsDir() {
				files = append(files, filepath.Join(dir, e.Name()))
			}
		}
	}
	return files, nil
}
