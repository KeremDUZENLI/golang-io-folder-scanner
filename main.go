package main

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/terminal"
)

func main() {
	cfg := env.ConfigDefault

	// 0_
	pathToScan := terminal.InputPath("[NEW] Path To scan")

	addFoldersToSkip := terminal.InputList("[ADD] Folders to skip", cfg.FoldersToSkip)
	allfoldersToSkip := append(cfg.FoldersToSkip, addFoldersToSkip...)

	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", cfg.SuffixesToScan)

	// 1_list
	foldersAll, err := scanner.ListFolders(pathToScan)
	terminal.PrintError("Failed listing foldersAll", err)

	// 2_filter
	foldersFiltered := scanner.FilterFoldersByName(foldersAll, allfoldersToSkip)

	filesFiltered, err := scanner.FilterFilesBySuffix(foldersFiltered, suffixesToScan)
	terminal.PrintError("Failed listing filesFiltered", err)

	// 3_content
	filesPathAndContent, err := scanner.ScanFilesContent(filesFiltered)
	terminal.PrintError("Failed listing filesPathAndContent", err)

	results := make([][2]string, 0, len(filesPathAndContent))
	for _, c := range filesPathAndContent {
		results = append(results, [2]string{c.Path, c.Content})
	}
	terminal.PrintScan(results)

	// 4_tree
	foldersFiltered2 := scanner.FilterFoldersByName(foldersAll, cfg.FoldersToSkip)

	filesAll, err := scanner.ListFiles(foldersFiltered2)
	terminal.PrintError("Failed listing filesAll", err)

	addFoldersTreeToSkip := terminal.InputList("[ADD] Folders tree to skip", cfg.FoldersTreeToSkip)
	allFoldersTreeToSkip := append(cfg.FoldersTreeToSkip, addFoldersTreeToSkip...)

	trees := scanner.CreateTree(pathToScan, foldersFiltered2, filesAll, allFoldersTreeToSkip)
	terminal.PrintTree(trees)

	// 5_find
	pathToScan2 := terminal.InputPath("[NEW] Path To scan")

	foldersAll2, err := scanner.ListFolders(pathToScan2)
	terminal.PrintError("Failed listing foldersAll", err)

	foldersEmpty := scanner.FindFoldersEmpty(foldersAll2)
	terminal.PrintFolders(foldersEmpty)

	// ----------------------

	pathToScan3 := terminal.InputPath("[NEW] Path To scan")

	foldersAll3, err := scanner.ListFolders(pathToScan3)
	terminal.PrintError("Failed listing foldersAll", err)

	suffixesToScan2 := terminal.InputList("[NEW] Suffixes to scan", cfg.SuffixesToScan)
	foldersByFileSuffix := scanner.FindFoldersByFileSuffix(foldersAll3, suffixesToScan2)
	terminal.PrintFolders(foldersByFileSuffix)

	terminal.InputKeypress()
}
