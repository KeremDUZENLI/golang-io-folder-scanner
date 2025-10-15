package scanner

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

// func PrintTree(cfg *env.Config, path, prefix string, skipFiles bool) error {
// 	if prefix == "" {
// 		fmt.Println("\nASCII_TREE=")
// 	}

// 	entries, err := os.ReadDir(path)
// 	if err != nil {
// 		return err
// 	}
// 	utils.SortEntries(entries)

// 	var filtered []os.DirEntry
// 	for _, e := range entries {
// 		if e.IsDir() && utils.Contains(cfg.Scan.FoldersToSkip, e.Name()) {
// 			continue
// 		}
// 		if skipFiles && !e.IsDir() {
// 			continue
// 		}
// 		filtered = append(filtered, e)
// 	}

// 	for i, e := range filtered {
// 		name := e.Name()
// 		fmt.Println(prefix + utils.TreeBranch(i, len(filtered)) + name)

// 		if !e.IsDir() {
// 			continue
// 		}

// 		nextSkip := skipFiles || utils.Contains(cfg.Tree.FoldersContentToSkip, name)
// 		if err := PrintTree(
// 			cfg,
// 			filepath.Join(path, name),
// 			prefix+utils.Indent(i, len(filtered)),
// 			nextSkip,
// 		); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func PrintTree(cfg *env.Config, rootPath string) error {
	return printTreeRecursive(cfg, rootPath, "", false)
}

func printTreeRecursive(cfg *env.Config, path, prefix string, skipFiles bool) error {
	if prefix == "" {
		fmt.Println("\nASCII_TREE=")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	utils.SortEntries(entries)

	var filtered []os.DirEntry
	for _, e := range entries {
		// Do NOT exclude directories based on cfg.Scan.FoldersToSkip here.
		// skipFiles controls whether files are printed at this level.
		if skipFiles && !e.IsDir() {
			continue
		}
		filtered = append(filtered, e)
	}

	for i, e := range filtered {
		name := e.Name()
		fmt.Println(prefix + utils.TreeBranch(i, len(filtered)) + name)

		if !e.IsDir() {
			continue
		}

		// If the directory name is listed in cfg.Tree.FoldersContentToSkip,
		// we print the directory but do NOT descend into it.
		nextSkip := skipFiles || utils.Contains(cfg.Tree.FoldersContentToSkip, name)
		if err := printTreeRecursive(
			cfg,
			filepath.Join(path, name),
			prefix+utils.Indent(i, len(filtered)),
			nextSkip,
		); err != nil {
			return err
		}
	}
	return nil
}
