package scanner

import (
	"io/fs"
	"path/filepath"
)

func FindFoldersEmpty(folders []string) []string {
	foldersEmpty := []string{}
	for _, root := range folders {
		hasFile := false
		filepath.WalkDir(root, func(p string, d fs.DirEntry, _ error) error {
			if d != nil && !d.IsDir() {
				hasFile = true
				return filepath.SkipDir
			}
			return nil
		})
		if !hasFile {
			foldersEmpty = append(foldersEmpty, root)
		}
	}
	return foldersEmpty
}

func FindFoldersByFileSuffix(folders, suffixesToFind []string) []string {
	foldersByFileSuffix := []string{}
	for _, root := range folders {
		found := false
		filepath.WalkDir(root, func(p string, d fs.DirEntry, _ error) error {
			if d == nil || d.IsDir() {
				return nil
			}
			if hasSuffix(p, suffixesToFind) {
				found = true
				return filepath.SkipDir
			}
			return nil
		})
		if found {
			foldersByFileSuffix = append(foldersByFileSuffix, root)
		}
	}
	return foldersByFileSuffix
}
