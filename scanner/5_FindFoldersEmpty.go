package scanner

import (
	"io/fs"
	"path/filepath"
)

func FindFoldersEmpty(folders []string) []string {
	foldersEmpty := []string{}
	for _, folder := range folders {
		hasFile, _ := walkFolderHas(folder, func(_ string) bool { return true })
		if !hasFile {
			foldersEmpty = append(foldersEmpty, folder)
		}
	}

	return foldersEmpty
}

func walkFolderHas(folder string, match func(filePath string) bool) (bool, error) {
	found := false
	err := filepath.WalkDir(folder, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d == nil || d.IsDir() {
			return nil
		}
		if match(p) {
			found = true
			return filepath.SkipDir
		}
		return nil
	})

	return found, err
}
