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

	lines := scanner.CreateTree(pathToScan, filteredFolders, filteredFiles, allFoldersTreeToSkip)
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
