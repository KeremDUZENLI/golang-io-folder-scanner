package main

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func main() {
	cfg := env.DefaultConfig

	cfg.GetForPath()

	// cfg.GetForScan()
	// scanner.PrintScan(&cfg, cfg.Path.PathToScan)

	// cfg.GetForTree()
	// scanner.PrintTree(&cfg, cfg.Path.PathToScan)

	cfg.GetForFolders()
	emptyFolders := scanner.FindEmptyFolders(&cfg, cfg.Path.PathToScan)
	scanner.PrintEmptyFolders(emptyFolders)

	utils.WaitForKeypress()
}
