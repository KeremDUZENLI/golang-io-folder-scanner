package scanner

import (
	"path/filepath"
	"strings"
)

func FilterFolders(folders, foldersToSkip []string) []string {
	if len(foldersToSkip) == 0 {
		return folders
	}

	foldersToSkipDict := make(map[string]struct{}, len(foldersToSkip))
	for _, folderToSkip := range foldersToSkip {
		foldersToSkipDict[strings.ToLower(folderToSkip)] = struct{}{}
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

func FilterFilesBySuffix(files, suffixesToScan []string) []string {
	if len(suffixesToScan) == 0 {
		return files
	}

	suffixesToScanList := make([]string, 0, len(suffixesToScan))
	for _, s := range suffixesToScan {
		suffixesToScanList = append(suffixesToScanList, strings.ToLower(s))
	}

	filteredFilesBySuffix := make([]string, 0, len(files))
	for _, file := range files {
		if hasSuffix(file, suffixesToScanList) {
			filteredFilesBySuffix = append(filteredFilesBySuffix, file)
		}
	}

	return filteredFilesBySuffix
}

func shouldSkipFolder(path string, foldersToSkip map[string]struct{}) bool {
	paths := strings.Split(path, "/")
	for _, folder := range paths {
		if _, ok := foldersToSkip[folder]; ok {
			return true
		}
	}

	return false
}

func hasSuffix(path string, suffixesToScan []string) bool {
	base := filepath.Base(path)
	for _, suffix := range suffixesToScan {
		if strings.HasSuffix(base, suffix) {
			return true
		}
	}

	return false
}
