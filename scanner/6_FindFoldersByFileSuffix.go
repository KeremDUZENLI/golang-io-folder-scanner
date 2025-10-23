package scanner

import "github.com/KeremDUZENLI/golang-io-folder-scanner/helper"

func FindFoldersByFileSuffix(folders []string, suffixesToScan []string) []string {
	foldersByFileSuffix := []string{}
	for _, folder := range folders {
		found, err := walkFolderHas(folder, func(p string) bool { return hasSuffixFile(p, suffixesToScan) })
		helper.PrintError(err)
		if found {
			foldersByFileSuffix = append(foldersByFileSuffix, folder)
		}
	}

	return foldersByFileSuffix
}
