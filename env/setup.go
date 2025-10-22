package env

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/terminal"
)

func (c *Config) RunPath() []string {
	c.PathToScan = terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	return scanner.ListFolders(c.PathToScan)
}

func (c *Config) RunScanFilesContent(folders []string) {
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	allfoldersToSkip := append(c.FoldersToSkip, foldersToSkip...)

	filteredFolders := scanner.FilterFolders(folders, allfoldersToSkip)
	filteredFoldersFiles := scanner.ListFiles(filteredFolders)

	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)
	filteredFiles := scanner.FilterFiles(filteredFoldersFiles, suffixesToScan)

	lines := scanner.ScanFilesContent(filteredFiles)
	terminal.PrintLines("CONTENT OF FILES", c.PathToScan, lines)
}

func (c *Config) RunTree(folders []string) {
	filteredFolders := scanner.FilterFolders(folders, c.FoldersToSkip)
	filteredFoldersFiles := scanner.ListFiles(filteredFolders)

	foldersTreeToSkip := terminal.InputList("[ADD] Folders tree to skip", c.FoldersTreeToSkip)
	allFoldersTreeToSkip := append(c.FoldersTreeToSkip, foldersTreeToSkip...)

	lines := scanner.CreateTree(filteredFolders, filteredFoldersFiles, allFoldersTreeToSkip)
	terminal.PrintLines("ASCII TREE", c.PathToScan, lines)
}

func (c *Config) RunFoldersEmpty(folders []string) {
	filteredFolders := scanner.FilterFolders(folders, c.FoldersToSkip)
	foldersEmpty := scanner.FindFoldersEmpty(filteredFolders)

	terminal.PrintLines("EMPTY FOLDERS", c.PathToScan, foldersEmpty)
}

func (c *Config) RunFoldersBySuffix(folders []string) {
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)
	filteredFolders := scanner.FilterFolders(folders, c.FoldersToSkip)
	foldersByFileSuffix := scanner.FindFoldersByFileSuffix(filteredFolders, suffixesToScan)

	terminal.PrintLines("FOUND FOLDERS", c.PathToScan, foldersByFileSuffix)
}

func (c *Config) RunFilesCompare(folders1, folders2 []string) {
	filteredFolders1 := scanner.FilterFolders(folders1, c.FoldersToSkip)
	filteredFolders2 := scanner.FilterFolders(folders2, c.FoldersToSkip)

	filteredFiles1 := scanner.ListFiles(filteredFolders1)
	filteredFiles2 := scanner.ListFiles(filteredFolders2)
	onlyIn1, onlyIn2 := scanner.CompareFiles(filteredFiles1, filteredFiles2)

	terminal.PrintCompare("FILE COMPARISON", folders1[0], folders2[0], onlyIn1, onlyIn2)
}
