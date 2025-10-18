package scanner

import (
	"path/filepath"
	"strings"
)

func FilterFoldersByName(folders, foldersToSkip []string) []string {
	if len(foldersToSkip) == 0 {
		return folders
	}
	filteredFoldersByName := make([]string, 0, len(folders))
	for _, folder := range folders {
		if skipFolder(folder, foldersToSkip) {
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
	for _, p := range files {
		if hasSuffix(p, suffixesToScan) {
			filteredFilesBySuffix = append(filteredFilesBySuffix, p)
		}
	}
	return filteredFilesBySuffix
}

func skipFolder(path string, foldersToSkip []string) bool {
	if len(foldersToSkip) == 0 {
		return false
	}
	base := strings.ToLower(filepath.Base(path))
	for _, folder := range foldersToSkip {
		if base == strings.ToLower(folder) {
			return true
		}
	}
	return false
}

func hasSuffix(path string, suffixesToScan []string) bool {
	if len(suffixesToScan) == 0 {
		return true
	}
	base := strings.ToLower(filepath.Base(path))
	for _, suffix := range suffixesToScan {
		if strings.HasSuffix(base, strings.ToLower(suffix)) {
			return true
		}
	}
	return false
}
