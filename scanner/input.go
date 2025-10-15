package scanner

import (
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func GetCurrentWorkingDirectory() (string, error) {
	return os.Getwd()
}

func FormatPathToScan(directoryToScan string) (string, error) {
	return filepath.Abs(directoryToScan)
}

func GetForPath(cfg *env.Config) {
	cwd, err := GetCurrentWorkingDirectory()
	utils.PrintError("Failed to Get Current Working Directory", err)

	input, err := utils.ReadInput("Path To Scan", cwd)
	utils.PrintError("Failed to Read Path To Scan", err)

	pathToScan, err := FormatPathToScan(input)
	utils.PrintError("Failed to Format Path to Scan", err)

	cfg.Path.PathToScan = pathToScan
}

func GetForScan(cfg *env.Config) {
	defaultSuffixesToScan := utils.ListToString(cfg.Scan.SuffixesToScan)
	input, err := utils.ReadInput("Suffixes to Scan", defaultSuffixesToScan)
	utils.PrintError("Failed to Read Suffixes to Scan", err)
	cfg.Scan.SuffixesToScan = utils.StringToList(input)

	defaultFoldersToSkip := utils.ListToString(cfg.Scan.FoldersToSkip)
	input, err = utils.ReadInput("Folders to Skip", defaultFoldersToSkip)
	utils.PrintError("Failed to Read Folders to Skip", err)
	input = defaultFoldersToSkip + "," + input
	cfg.Scan.FoldersToSkip = utils.StringToList(input)
}

func GetForTree(cfg *env.Config) {
	DefaultFoldersContentToSkip := utils.ListToString(cfg.Tree.FoldersContentToSkip)
	input, err := utils.ReadInput("Folders Content to Skip", DefaultFoldersContentToSkip)
	utils.PrintError("Failed to Read Folders Content to Skip", err)
	cfg.Tree.FoldersContentToSkip = utils.StringToList(input)
}
