package main

import (
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func main() {
	cfg := env.ConfigDefault
	cfg.PathToScan, _ = os.Getwd()
	folders := cfg.RunPath()

	cfg.RunTester(folders, "0.1, 0.2")

	// cfg.RunScanFilesContent(folders)
	// cfg.RunTree(folders)
	// cfg.RunFoldersEmpty(folders)
	// cfg.RunFoldersBySuffix(folders)

	// foldersCompare := cfg.RunPath()
	// cfg.RunFilesCompare(folders, foldersCompare)
}
