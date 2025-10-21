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
	for _, s := range foldersToSkip {
		foldersToSkipDict[strings.ToLower(s)] = struct{}{}
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

	filteredFilesBySuffix := make([]string, 0, len(files))
	for _, file := range files {
		if hasSuffix(file, suffixesToScan) {
			filteredFilesBySuffix = append(filteredFilesBySuffix, file)
		}
	}

	return filteredFilesBySuffix
}

func shouldSkipFolder(path string, foldersToSkip map[string]struct{}) bool {
	paths := strings.Split(filepath.ToSlash(path), "/")
	for _, folder := range paths {
		if _, ok := foldersToSkip[strings.ToLower(folder)]; ok {
			return true
		}
	}

	return false
}

func hasSuffix(path string, suffixesToScan []string) bool {
	base := strings.ToLower(filepath.Base(path))
	for _, suffix := range suffixesToScan {
		if strings.HasSuffix(base, strings.ToLower(suffix)) {
			return true
		}
	}

	return false
}
