package main

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func main() {
	cfg := env.ConfigDefault
	cfgAdd := env.ConfigAdd{}

	inputPathToScan := utils.ReadInputPath("[OVERRIDE] Path To Scan", cfg.PathToScan)
	cfg.PathToScan = inputPathToScan

	inputSuffixesToScan := utils.ReadInputList("[OVERRIDE] Suffixes to Scan", cfg.SuffixesToScan)
	cfg.SuffixesToScan = inputSuffixesToScan

	inputFoldersToSkip := utils.ReadInputList("[ADD] Folders to Skip", cfg.FoldersToSkip)
	cfgAdd.FoldersToSkip = inputFoldersToSkip

	inputFoldersContentToSkip := utils.ReadInputList("[ADD] Folders Content to Skip", cfg.FoldersContentToSkip)
	cfgAdd.FoldersContentToSkip = inputFoldersContentToSkip

	scan := scanner.ScanFiles(&cfg, &cfgAdd)
	utils.PrintScan(scan)

	trees := scanner.GetTrees(&cfg, &cfgAdd)
	utils.PrintTree(trees)

	emptyFolders := scanner.FindEmptyFolders(&cfg)
	utils.PrintEmptyFolders(emptyFolders)

	utils.WaitForKeypress()
}
