package env

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/terminal"
)

func (c config) RunPath() []string {
	c.PathToScan = terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	return scanner.ListFolders(c.PathToScan)
}

func (c config) RunScanFilesContent(folders []string) {
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)

	filteredFoldersByName := scanner.FilterFoldersByName(folders, allfoldersToSkip)
	filteredFilesByName := scanner.ListFiles(filteredFoldersByName)

	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)
	filteredFilesBySuffix := scanner.FilterFilesBySuffix(filteredFilesByName, suffixesToScan)

	lines := scanner.ScanFilesContent(filteredFilesBySuffix)
	terminal.PrintLines("CONTENT OF FILES", lines)
}

func (c config) RunTree(folders []string) {
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	foldersTreeToSkip := terminal.InputList("[ADD] Folders tree to skip", c.FoldersTreeToSkip)

	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)
	allFoldersTreeToSkip := append(c.FoldersTreeToSkip, foldersTreeToSkip...)

	filteredFolders := scanner.FilterFoldersByName(folders, allfoldersToSkip)
	filteredFiles := scanner.ListFiles(filteredFolders)

	lines := scanner.CreateTree(filteredFolders, filteredFiles, allFoldersTreeToSkip)
	terminal.PrintLines("ASCII TREE", lines)
}

func (c config) RunFoldersEmpty(folders []string) {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)

	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)

	filteredFolders := scanner.FilterFoldersByName(folders, allfoldersToSkip)
	foldersEmpty := scanner.FindFoldersEmpty(filteredFolders)

	terminal.PrintFolders("EMPTY FOLDERS", pathToScan, foldersEmpty)
}

func (c config) RunFoldersBySuffix(folders []string) {
	pathToScan := terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)

	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)

	filteredFolders := scanner.FilterFoldersByName(folders, allfoldersToSkip)
	foldersByFileSuffix := scanner.FindFoldersByFileSuffix(filteredFolders, suffixesToScan)

	terminal.PrintFolders("FOUND FOLDERS", pathToScan, foldersByFileSuffix)
}

func (c config) RunFilesCompare(folders1, folders2 []string) {
	filteredFolders1 := scanner.FilterFoldersByName(folders1, c.FoldersToSkip)
	filteredFolders2 := scanner.FilterFoldersByName(folders2, c.FoldersToSkip)

	filteredFiles1 := scanner.ListFiles(filteredFolders1)
	filteredFiles2 := scanner.ListFiles(filteredFolders2)
	onlyIn1, onlyIn2 := scanner.CompareFiles(filteredFiles1, filteredFiles2)

	terminal.PrintCompare("FILE COMPARISON", folders1[0], folders2[0], onlyIn1, onlyIn2)
}
