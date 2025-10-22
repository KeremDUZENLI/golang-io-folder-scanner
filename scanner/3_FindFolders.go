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

func FindFoldersByFileSuffix(folders []string, suffixesToFind []string) []string {
	foldersByFileSuffix := []string{}
	for _, folder := range folders {
		found, err := walkFolderHas(folder, func(p string) bool { return hasSuffixFile(p, suffixesToFind) })
		if err != nil {
			return nil
		}
		if found {
			foldersByFileSuffix = append(foldersByFileSuffix, folder)
		}
	}

	return foldersByFileSuffix
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
