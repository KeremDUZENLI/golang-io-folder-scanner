package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func PrintTree(path string, cfg *env.Config) error {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nASCII_TREE=")
	return printTreeRecursive(path, cfg, "", false)
}

func printTreeRecursive(path string, cfg *env.Config, prefix string, skipFiles bool) error {
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
			err := printTreeRecursive(filepath.Join(path, name), cfg, prefix+indent(i, len(filtered)), nextSkip)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
