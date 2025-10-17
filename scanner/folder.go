package scanner

import (
	"path/filepath"
	"strings"
)

func FindEmptyFoldersFromLists(path string, folders, files []string) []string {
	nonEmpty := make(map[string]bool, len(folders))

	for _, f := range files {
		dir := filepath.Dir(f)
		for {
			nonEmpty[dir] = true
			if dir == path {
				break
			}

			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}

			dir = parent
		}
	}

	var empty []string
	for _, folder := range folders {
		rel, err := filepath.Rel(path, folder)

		if err != nil {
			continue
		}

		if strings.HasPrefix(rel, "..") {
			continue
		}

		if !nonEmpty[folder] {
			empty = append(empty, folder)
		}
	}

	return empty
}
