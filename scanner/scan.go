package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func PrintScan(cfg *env.Config, path string) error {
	return printScanRecursive(cfg, path)
}

func printScanRecursive(cfg *env.Config, path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	sortEntries(entries)

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(path, name)

		if entry.IsDir() {
			if contains(cfg.Scan.FolderToSkip, name) {
				continue
			}
			if err := printScanRecursive(cfg, fullPath); err != nil {
				return err
			}
			continue
		}

		if hasSuffix(name, cfg.Scan.SuffixesToScan) {
			data, err := os.ReadFile(fullPath)
			if err != nil {
				return err
			}

			relPath, _ := filepath.Rel(cfg.Path.PathToScan, fullPath)
			fmt.Printf("\n%s=\n%s\n", relPath, data)
			fmt.Println(strings.Repeat("-", 100))
		}
	}

	return nil
}
