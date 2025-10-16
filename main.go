package main

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func main() {
	cfg := env.DefaultConfig
	cfg.GetForPath()

	cfg.GetForScan()
	scanner.PrintScan(cfg.Path.PathToScan, &cfg)

	cfg.GetForTree()
	scanner.PrintTree(cfg.Path.PathToScan, &cfg)

	cfg.GetForFolders()
	scanner.PrintEmptyFolders(cfg.Path.PathToScan)

	utils.WaitForKeypress()
}
