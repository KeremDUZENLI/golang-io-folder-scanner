package scanner

func FindFoldersByFileSuffix(folders []string, suffixesToScan []string) []string {
	foldersByFileSuffix := []string{}
	for _, folder := range folders {
		found, err := walkFolderHas(folder, func(p string) bool { return hasSuffixFile(p, suffixesToScan) })
		if err != nil {
			return nil
		}
		if found {
			foldersByFileSuffix = append(foldersByFileSuffix, folder)
		}
	}

	return foldersByFileSuffix
}
