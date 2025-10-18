package main

import (
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func main() {
	cfg := env.ConfigDefault
	cfg.PathToScan, _ = os.Getwd()

	cfg.RunFilesPathAndContent()
	// cfg.RunAsciiTree()
	// cfg.RunFindFoldersEmpty()
	// cfg.RunFindFoldersBySuffix()
}
