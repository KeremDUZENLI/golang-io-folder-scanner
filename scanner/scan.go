package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func PrintScan(path string, cfg *env.Config) error {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nSCANNED_FILES=")
	return printScanRecursive(path, cfg)
}

func printScanRecursive(path string, cfg *env.Config) error {
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
			if err := printScanRecursive(fullPath, cfg); err != nil {
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
