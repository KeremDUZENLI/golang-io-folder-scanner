package utils

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func GetForPath(cfg *env.Config) {
	cwd, err := GetCurrentWorkingDirectory()
	PrintError("Failed to Get Current Working Directory", err)

	input, err := ReadInput("[OVERRIDE] Path To Scan", cwd)
	PrintError("Failed to Read Path To Scan", err)

	pathToScan, err := FormatPathToScan(input)
	PrintError("Failed to Format Path to Scan", err)

	cfg.Path.PathToScan = pathToScan
}

func GetForScan(cfg *env.Config) {
	defaultSuffixesToScan := ListToString(cfg.Scan.SuffixesToScan)
	input, err := ReadInput("[OVERRIDE] Suffixes to Scan", defaultSuffixesToScan)
	PrintError("Failed to Read Suffixes to Scan", err)
	cfg.Scan.SuffixesToScan = StringToList(input)

	defaultFoldersToSkip := ListToString(cfg.Scan.DefaultFoldersToSkip)
	input, err = ReadInput("[ADD] Folders to Skip", defaultFoldersToSkip)
	PrintError("Failed to Read Folders to Skip", err)
	cfg.Scan.FolderToSkip = StringToList(defaultFoldersToSkip + "," + input)
}

func GetForTree(cfg *env.Config) {
	defaultFoldersContentToSkip := ListToString(cfg.Tree.DefaultFoldersContentToSkip)
	input, err := ReadInput("[ADD] Folders Content to Skip", defaultFoldersContentToSkip)
	PrintError("Failed to Read Folders Content to Skip", err)
	cfg.Tree.FoldersContentToSkip = StringToList(defaultFoldersContentToSkip + "," + input)
}
