package env

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func (c *Config) GetForPath() {
	cwd, err := getCurrentWorkingDirectory()
	utils.PrintError("Failed to Get Current Working Directory", err)

	input, err := readInput("[OVERRIDE] Path To Scan", cwd)
	utils.PrintError("Failed to Read Path To Scan", err)

	pathToScan, err := formatPathToScan(input)
	utils.PrintError("Failed to Format Path to Scan", err)

	c.Path.PathToScan = pathToScan
}

func (c *Config) GetForScan() {
	defaultSuffixesToScan := listToString(c.Scan.SuffixesToScan)
	input, err := readInput("[OVERRIDE] Suffixes to Scan", defaultSuffixesToScan)
	utils.PrintError("Failed to Read Suffixes to Scan", err)
	c.Scan.SuffixesToScan = stringToList(input)

	defaultFoldersToSkip := listToString(c.Scan.DefaultFoldersToSkip)
	input, err = readInput("[ADD] Folders to Skip", defaultFoldersToSkip)
	utils.PrintError("Failed to Read Folders to Skip", err)
	c.Scan.FolderToSkip = stringToList(defaultFoldersToSkip + "," + input)
}

func (c *Config) GetForTree() {
	defaultFoldersContentToSkip := listToString(c.Tree.DefaultFoldersContentToSkip)
	input, err := readInput("[ADD] Folders Content to Skip", defaultFoldersContentToSkip)
	utils.PrintError("Failed to Read Folders Content to Skip", err)
	c.Tree.FoldersContentToSkip = stringToList(defaultFoldersContentToSkip + "," + input)
}

func (c *Config) GetForFolders() {
	folderToScan, err := getFoldersToScan(c.Path.PathToScan)
	utils.PrintError("Failed to Read Directory", err)
	c.Folders.FoldersToScan = folderToScan
}
