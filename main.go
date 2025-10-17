package main

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func main() {
	cfg := env.ConfigDefault

	// override
	cfg.PathToScan = utils.ReadInputPath("[OVERRIDE] Path To Scan")
	cfg.SuffixesToScan = utils.ReadInputList("[OVERRIDE] Suffixes to Scan", cfg.SuffixesToScan)

	// add
	extraFoldersToSkip := utils.ReadInputList("[ADD] Folders to Skip", cfg.FoldersToSkip)
	extraFoldersContentToSkip := utils.ReadInputList("[ADD] Folders Content to Skip", cfg.FoldersContentToSkip)

	// 0) list folders + files (absolute)
	foldersToSkip := append(cfg.FoldersToSkip, extraFoldersToSkip...)

	folders, err := scanner.ListFolders(cfg.PathToScan, foldersToSkip)
	utils.PrintError("Failed listing folders", err)

	files, err := scanner.ListFiles(cfg.PathToScan, cfg.SuffixesToScan, foldersToSkip)
	utils.PrintError("Failed listing files", err)

	// 1) scan files
	fileResults, _ := scanner.ReadFiles(files)
	results := make([][2]string, 0, len(fileResults))
	for _, fr := range fileResults {
		results = append(results, [2]string{fr.Path, fr.Content})
	}
	utils.PrintScan(results)

	// 2) ascii tree
	foldersContentToSkip := append(cfg.FoldersContentToSkip, extraFoldersContentToSkip...)

	tree := scanner.GetTreesFromLists(cfg.PathToScan, folders, files, foldersContentToSkip)
	utils.PrintTree(tree)

	// 3) empty folders
	empty := scanner.FindEmptyFoldersFromLists(cfg.PathToScan, folders, files)
	utils.PrintEmptyFolders(empty)

	utils.WaitForKeypress()
}
