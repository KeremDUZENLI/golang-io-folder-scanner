package scanner

import (
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func GetTrees(cfg *env.Config) []string {
	return getTreesRecursive(cfg.Path.PathToScan, cfg.Scan.DefaultFoldersToSkip, cfg.Tree.FoldersContentToSkip, false, "")
}

func getTreesRecursive(path string, defaultFoldersToSkip, foldersContentToSkip []string, skipFiles bool, prefix string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		env.PrintError("Failed Reading Directory", err)
		return nil
	}
	sortEntries(entries)

	var filtered []os.DirEntry
	for _, e := range entries {
		if e.IsDir() && contain(defaultFoldersToSkip, e.Name()) {
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
			childPath := filepath.Join(path, name)
			childLines := getTreesRecursive(childPath, defaultFoldersToSkip, foldersContentToSkip, nextSkip, prefix+indent(i, len(filtered)))
			trees = append(trees, childLines...)
		}
	}
	return trees
}
