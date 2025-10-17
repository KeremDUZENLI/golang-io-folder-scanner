package scanner

import (
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func GetTrees(cfg *env.Config, cfgAdd *env.ConfigAdd) []string {
	foldersContentToSkipTotal := append(cfg.FoldersContentToSkip, cfgAdd.FoldersContentToSkip...)
	return getTreesRecursive(cfg.PathToScan, cfg.FoldersToSkip, foldersContentToSkipTotal, false, "")
}

func getTreesRecursive(pathToScan string, foldersToSkip, foldersContentToSkip []string, skipFiles bool, prefix string) []string {
	entries, err := os.ReadDir(pathToScan)
	if err != nil {
		return nil
	}
	sortEntries(entries)

	var filtered []os.DirEntry
	for _, e := range entries {
		if e.IsDir() && contain(foldersToSkip, e.Name()) {
			continue
		}
		if skipFiles && !e.IsDir() {
			continue
		}
		filtered = append(filtered, e)
	}

	var trees []string
	for i, e := range filtered {
		name := e.Name()
		line := prefix + treeBranch(i, len(filtered)) + name
		trees = append(trees, line)

		if e.IsDir() {
			nextSkip := skipFiles || contain(foldersContentToSkip, name)
			childPath := filepath.Join(pathToScan, name)
			childLines := getTreesRecursive(childPath, foldersToSkip, foldersContentToSkip, nextSkip, prefix+indent(i, len(filtered)))
			trees = append(trees, childLines...)
		}
	}
	return trees
}
