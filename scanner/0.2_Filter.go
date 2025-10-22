package scanner

import (
	"path/filepath"
	"strings"
)

func FilterFolders(folders, foldersToSkip []string) []string {
	foldersToSkipDict := make(map[string]struct{}, len(foldersToSkip))
	for _, folderToSkip := range foldersToSkip {
		foldersToSkipDict[folderToSkip] = struct{}{}
	}

	filteredFoldersByName := make([]string, 0, len(folders))
	for _, folder := range folders {
		if shouldSkipFolder(folder, foldersToSkipDict) {
			continue
		}
		filteredFoldersByName = append(filteredFoldersByName, folder)
	}

	return filteredFoldersByName
}

func FilterFiles(files, suffixesToScan []string) []string {
	filteredFilesBySuffix := make([]string, 0, len(files))
	for _, file := range files {
		if hasSuffixFile(file, suffixesToScan) {
			filteredFilesBySuffix = append(filteredFilesBySuffix, file)
		}
	}

	return filteredFilesBySuffix
}

func shouldSkipFolder(folder string, foldersToSkip map[string]struct{}) bool {
	paths := strings.Split(folder, "/")
	for _, folder := range paths {
		if _, ok := foldersToSkip[folder]; ok {
			return true
		}
	}

	return false
}

func hasSuffixFile(file string, suffixesToScan []string) bool {
	base := filepath.Base(file)
	for _, suffix := range suffixesToScan {
		if strings.HasSuffix(base, suffix) {
			return true
		}
	}

	return false
}
