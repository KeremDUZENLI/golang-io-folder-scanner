package scanner

import (
	"io/fs"
	"path/filepath"
)

func FindFoldersEmpty(folders []string) ([]string, error) {
	foldersEmpty := []string{}
	for _, folder := range folders {
		hasFile, err := walkFolderHas(folder, func(_ string) bool { return true })
		if err != nil {
			return nil, err
		}
		if !hasFile {
			foldersEmpty = append(foldersEmpty, folder)
		}
	}
	return foldersEmpty, nil
}

func FindFoldersByFileSuffix(folders []string, suffixesToFind []string) ([]string, error) {
	foldersByFileSuffix := []string{}
	for _, folder := range folders {
		found, err := walkFolderHas(folder, func(p string) bool { return hasSuffix(p, suffixesToFind) })
		if err != nil {
			return nil, err
		}
		if found {
			foldersByFileSuffix = append(foldersByFileSuffix, folder)
		}
	}
	return foldersByFileSuffix, nil
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
