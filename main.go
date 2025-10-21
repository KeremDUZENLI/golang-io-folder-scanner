package main

import (
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func main() {
	cfg := env.ConfigDefault
	cfg.PathToScan, _ = os.Getwd()
	folders := cfg.RunPath()

	// folders = scanner.FilterFolders(folders, cfg.FoldersToSkip)
	// for _, i := range folders {
	// 	println(i)
	// }

	// println("***************************")

	// files := scanner.ListFiles(folders)
	// for _, i := range files {
	// 	println(i)
	// }

	cfg.RunScanFilesContent(folders)
	cfg.RunTree(folders)
	cfg.RunFoldersEmpty(folders)
	cfg.RunFoldersBySuffix(folders)

	foldersCompare := cfg.RunPath()
	cfg.RunFilesCompare(folders, foldersCompare)
}
