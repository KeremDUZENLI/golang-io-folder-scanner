package env

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/terminal"
)

func (c Config) RunFilesContent() {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)

	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)

	foldersAll, err := scanner.ListFolders(pathToScan)
	terminal.PrintError("Failed listing foldersAll", err)

	filteredFoldersByName := scanner.FilterFoldersByName(foldersAll, allfoldersToSkip)
	filteredFilesByName, err := scanner.ListFiles(filteredFoldersByName)
	terminal.PrintError("Failed listing filteredFilesByName", err)

	filteredFilesBySuffix := scanner.FilterFilesBySuffix(filteredFilesByName, suffixesToScan)
	lines := scanner.ScanFilesContent(filteredFilesBySuffix)
	terminal.PrintLines("CONTENT OF FILES", lines)
}

func (c Config) RunTree() {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	foldersTreeToSkip := terminal.InputList("[ADD] Folders tree to skip", c.FoldersTreeToSkip)

	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)
	allFoldersTreeToSkip := append(c.FoldersTreeToSkip, foldersTreeToSkip...)

	foldersAll, err := scanner.ListFolders(pathToScan)
	terminal.PrintError("Failed listing foldersAll", err)

	filteredFolders := scanner.FilterFoldersByName(foldersAll, allfoldersToSkip)
	filteredFiles, err := scanner.ListFiles(filteredFolders)
	terminal.PrintError("Failed listing filteredFiles", err)

	lines := scanner.CreateTree(filteredFolders, filteredFiles, allFoldersTreeToSkip)
	terminal.PrintLines("ASCII TREE", lines)
}

func (c Config) RunFoldersEmpty() {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)

	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)

	foldersAll, err := scanner.ListFolders(pathToScan)
	terminal.PrintError("Failed listing foldersAll", err)

	filteredFolders := scanner.FilterFoldersByName(foldersAll, allfoldersToSkip)
	foldersEmpty, err := scanner.FindFoldersEmpty(filteredFolders)
	terminal.PrintError("Failed listing foldersEmpty", err)

	terminal.PrintFolders("EMPTY FOLDERS", pathToScan, foldersEmpty)
}

func (c Config) RunFoldersBySuffix() {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)

	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)

	foldersAll, err := scanner.ListFolders(pathToScan)
	terminal.PrintError("Failed listing foldersAll", err)

	filteredFolders := scanner.FilterFoldersByName(foldersAll, allfoldersToSkip)
	foldersByFileSuffix, err := scanner.FindFoldersByFileSuffix(filteredFolders, suffixesToScan)
	terminal.PrintError("Failed listing folderByFileSuffix", err)

	terminal.PrintFolders("FOUND FOLDERS", pathToScan, foldersByFileSuffix)
}

func (c Config) RunFilesCompare() {
	pathToScan1 := terminal.InputPath("[NEW] Path to scan 1", c.PathToScan)
	pathToScan2 := terminal.InputPath("[NEW] Path to scan 2", c.PathToScan)

	foldersAll1, err := scanner.ListFolders(pathToScan1)
	terminal.PrintError("Failed listing foldersAll1", err)
	foldersAll2, err := scanner.ListFolders(pathToScan2)
	terminal.PrintError("Failed listing foldersAll1", err)

	filteredFolders1 := scanner.FilterFoldersByName(foldersAll1, c.FoldersToSkip)
	filteredFolders2 := scanner.FilterFoldersByName(foldersAll2, c.FoldersToSkip)

	filteredFiles1, err := scanner.ListFiles(filteredFolders1)
	terminal.PrintError("Failed listing filteredFiles1", err)
	filteredFiles2, err := scanner.ListFiles(filteredFolders2)
	terminal.PrintError("Failed listing filteredFiles2", err)

	onlyIn1, onlyIn2, err := scanner.CompareFiles(filteredFiles1, filteredFiles2)
	terminal.PrintError("Failed listing differences", err)

	terminal.PrintCompare("FILE COMPARISON", pathToScan1, pathToScan2, onlyIn1, onlyIn2)
}
