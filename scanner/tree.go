package scanner

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func PrintTree(cfg *env.Config, path string) error {
	fmt.Println("\nASCII_TREE=")
	return printTreeRecursive(cfg, path, "", false)
}

func printTreeRecursive(cfg *env.Config, path, prefix string, skipFiles bool) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	sortEntries(entries)

	var filtered []os.DirEntry
	for _, e := range entries {
		if e.IsDir() && contains(cfg.Scan.DefaultFoldersToSkip, e.Name()) {
			continue
		}
		if skipFiles && !e.IsDir() {
			continue
		}
		filtered = append(filtered, e)
	}

	for i, e := range filtered {
		name := e.Name()
		fmt.Println(prefix + treeBranch(i, len(filtered)) + name)

		if e.IsDir() {
			nextSkip := skipFiles || contains(cfg.Tree.FoldersContentToSkip, name)
			err := printTreeRecursive(cfg, filepath.Join(path, name), prefix+indent(i, len(filtered)), nextSkip)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
