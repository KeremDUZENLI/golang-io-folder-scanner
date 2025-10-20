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
	filesPathAndContent := scanner.ScanFilesContent(filteredFilesBySuffix)
	terminal.PrintError("Failed listing filesPathAndContent", err)

	terminal.PrintLines("CONTENT OF FILES", filesPathAndContent)
}

func (c Config) RunTree() {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	foldersTreeToSkip := terminal.InputList("[ADD] Folders tree to skip", c.FoldersTreeToSkip)

	allFoldersTreeToSkip := append(c.FoldersTreeToSkip, foldersTreeToSkip...)

	foldersAll, err := scanner.ListFolders(pathToScan)
	terminal.PrintError("Failed listing foldersAll", err)

	filteredFolders := scanner.FilterFoldersByName(foldersAll, c.FoldersToSkip)
	filteredFiles, err := scanner.ListFiles(filteredFolders)
	terminal.PrintError("Failed listing filteredFiles", err)

	tree := scanner.CreateTree(pathToScan, filteredFolders, filteredFiles, allFoldersTreeToSkip)

	terminal.PrintLines("ASCII TREE", tree)
}

func (c Config) RunFoldersEmpty() {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)

	foldersAll, err := scanner.ListFolders(pathToScan)
	terminal.PrintError("Failed listing foldersAll", err)

	filteredFolders := scanner.FilterFoldersByName(foldersAll, c.FoldersToSkip)
	foldersEmpty, err := scanner.FindFoldersEmpty(filteredFolders)
	terminal.PrintError("Failed listing foldersEmpty", err)

	terminal.PrintFolders("EMPTY FOLDERS", pathToScan, foldersEmpty)
}

func (c Config) RunFoldersBySuffix() {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)

	foldersAll, err := scanner.ListFolders(pathToScan)
	terminal.PrintError("Failed listing foldersAll", err)

	filteredFolders := scanner.FilterFoldersByName(foldersAll, c.FoldersToSkip)
	foldersByFileSuffix, err := scanner.FindFoldersByFileSuffix(filteredFolders, suffixesToScan)
	terminal.PrintError("Failed listing folderByFileSuffix", err)

	terminal.PrintFolders("FOUND FOLDERS", pathToScan, foldersByFileSuffix)
}
