package scanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

var (
	SuffixesToScan     = []string{".go"}
	SkipFolders        = []string{"__pycache__", ".venv", ".git", "_scripts"}
	SkipFoldersContent = []string{"data"}
)

func HandleFile(path string, d fs.DirEntry) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s=\n%s\n", path, data)
	fmt.Println(strings.Repeat("-", 100))
	return nil
}

func Traverse(path string, handle func(string, fs.DirEntry) error) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	utils.SortEntries(entries)

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(path, name)

		if entry.IsDir() {
			if utils.Contains(SkipFolders, name) || utils.Contains(SkipFoldersContent, name) {
				continue
			}
			if err := Traverse(fullPath, handle); err != nil {
				return err
			}
			continue
		}

		if utils.HasAnySuffix(name, SuffixesToScan) {
			if utils.Contains(SkipFoldersContent, filepath.Base(filepath.Dir(fullPath))) {
				continue
			}
			if err := handle(fullPath, entry); err != nil {
				return err
			}
		}
	}

	return nil
}
