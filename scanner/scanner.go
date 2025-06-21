package scanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

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
	return nil
}

func Traverse(root string, handle func(string, fs.DirEntry) error) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		name := d.Name()

		if d.IsDir() {
			if utils.Contains(SkipFolders, name) || utils.Contains(SkipFoldersContent, name) {
				return filepath.SkipDir
			}
		}

		if !d.IsDir() && utils.HasAnySuffix(name, SuffixesToScan) {
			if utils.Contains(SkipFoldersContent, filepath.Base(filepath.Dir(path))) {
				return nil
			}
			return handle(path, d)
		}
		return nil
	})
}
