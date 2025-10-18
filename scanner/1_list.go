package scanner

import (
	"io/fs"
	"os"
	"path/filepath"
)

func ListFolders(path string) ([]string, error) {
	folders := []string{}
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			folders = append(folders, p)
		}
		return nil
	})
	return folders, err
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
