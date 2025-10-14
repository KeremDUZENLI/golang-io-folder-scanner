package scanner

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func GetForPath(cfg *env.Config) {
	cwd, err := utils.GetCurrentWorkingDirectory()
	utils.PrintError("Failed to get current working directory", err)

	utils.PrintPrompt("Path To Scan", cwd)
	directoryToScan := utils.ReadInput(cwd)

	abs_path, err := utils.ResolveAbsolutePath(directoryToScan)
	utils.PrintError("Failed to resolve", err)

	cfg.Path.PathToScan = abs_path
}

func GetForScan(cfg *env.Config) {
	utils.PrintPrompt("Suffixes To Scan", utils.JoinStrings(cfg.Scan.SuffixesToScan))
	suffixesToScan := utils.ReadInput(utils.JoinStrings(cfg.Scan.SuffixesToScan))
	cfg.Scan.SuffixesToScan = utils.UpdateListIfInput(suffixesToScan)

	utils.PrintPrompt("Skip Folders", utils.JoinStrings(cfg.Scan.FoldersToSkip))
	foldersToSkip := utils.ReadInput(utils.JoinStrings(cfg.Scan.FoldersToSkip))
	cfg.Scan.FoldersToSkip = utils.UpdateListIfInput(foldersToSkip)
}

func GetForTree(cfg *env.Config) {
	utils.PrintPrompt("Skip Folders Content", utils.JoinStrings(cfg.Tree.FoldersContentToSkip))
	folderContentToSkip := utils.ReadInput(utils.JoinStrings(cfg.Tree.FoldersContentToSkip))
	cfg.Tree.FoldersContentToSkip = utils.UpdateListIfInput(folderContentToSkip)
}
