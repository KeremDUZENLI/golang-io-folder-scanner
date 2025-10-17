package scanner

func FilterFoldersByName(folders []string, foldersToSkip []string) []string {
	if len(foldersToSkip) == 0 {
		return folders
	}

	filteredFolders := make([]string, 0, len(folders))
	for _, folder := range folders {
		if skipFolder(folder, foldersToSkip) {
			continue
		}
		filteredFolders = append(filteredFolders, folder)
	}

	return filteredFolders
}

func FilterFilesBySuffix(files []string, suffixesToScan []string) []string {
	if len(suffixesToScan) == 0 {
		return files
	}

	filteredFiles := make([]string, 0, len(files))
	for _, file := range files {
		if hasSuffix(file, suffixesToScan) {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles
}
