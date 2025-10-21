package scanner

import (
	"os"
	"path/filepath"
)

func ListFolders(root string) []string {
	root = canonicalPath(root)

	folders := make([]string, 0, 64)
	stack := []string{root}

	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		folders = append(folders, dir)

		fd, err := os.Open(dir)
		if err != nil {
			continue
		}
		defer fd.Close()

		entries, err := fd.ReadDir(-1) // DO NOT sort; native order defines traversal & file order
		if err != nil {
			continue
		}

		childDirs := make([]string, 0, 8)
		for _, e := range entries {
			if e.IsDir() {
				childDirs = append(childDirs, canonicalPath(filepath.Join(dir, e.Name())))
			}
		}
		for i := len(childDirs) - 1; i >= 0; i-- {
			stack = append(stack, childDirs[i])
		}
	}

	return folders
}

func canonicalPath(path string) string {
	pathAbs, _ := filepath.Abs(path)
	clean := filepath.Clean(pathAbs)
	return filepath.ToSlash(clean)
}
