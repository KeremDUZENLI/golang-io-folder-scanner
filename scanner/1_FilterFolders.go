package scanner

import "strings"

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

func shouldSkipFolder(folder string, foldersToSkip map[string]struct{}) bool {
	paths := strings.Split(folder, "/")
	for _, folder := range paths {
		if _, ok := foldersToSkip[folder]; ok {
			return true
		}
	}

	return false
}
