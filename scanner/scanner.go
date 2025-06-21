package scanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func HandleFile(cfg *env.Config, path string, d fs.DirEntry) error {
	relPath, err := filepath.Rel(cfg.ScanRoot, path)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fmt.Printf("\n%s=\n%s\n", relPath, data)
	fmt.Println(strings.Repeat("-", 100))
	return nil
}

func Traverse(cfg *env.Config, path string, handle func(*env.Config, string, fs.DirEntry) error) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	utils.SortEntries(entries)

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(path, name)

		if entry.IsDir() {
			if utils.Contains(cfg.SkipFolders, name) || utils.Contains(cfg.SkipFoldersContent, name) {
				continue
			}
			if err := Traverse(cfg, fullPath, handle); err != nil {
				return err
			}
			continue
		}

		if utils.HasAnySuffix(name, cfg.SuffixesToScan) {
			if utils.Contains(cfg.SkipFoldersContent, filepath.Base(filepath.Dir(fullPath))) {
				continue
			}
			if err := handle(cfg, fullPath, entry); err != nil {
				return err
			}
		}
	}

	return nil
}
