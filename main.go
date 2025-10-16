package main

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func main() {
	cfg := env.DefaultConfig

	cfg.SetForPath()

	cfg.SetForScan()
	scan := scanner.ScanFiles(&cfg)
	utils.PrintScan(scan)

	cfg.SetForTree()
	trees := scanner.GetTrees(&cfg)
	utils.PrintTree(trees)

	cfg.SetForFolders()
	emptyFolders := scanner.FindEmptyFolders(&cfg)
	utils.PrintEmptyFolders(emptyFolders)

	utils.WaitForKeypress()
}
