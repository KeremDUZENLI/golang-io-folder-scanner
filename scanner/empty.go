package scanner

import (
	"fmt"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func FindEmptyFolders(cfg *env.Config, path string) []string {
	var emptyFolders []string

	for _, entry := range cfg.Folders.FoldersToScan {
		if entry.IsDir() {
			dirPath := filepath.Join(path, entry.Name())

			if !hasFiles(dirPath) {
				emptyFolders = append(emptyFolders, dirPath)
			}

			subEmptyFolders := FindEmptyFolders(cfg, dirPath)
			emptyFolders = append(emptyFolders, subEmptyFolders...)
		}
	}

	return emptyFolders
}

func PrintEmptyFolders(emptyFolders []string) {
	if len(emptyFolders) == 0 {
		fmt.Println("No empty folders found.")
		return
	}

	fmt.Println("Empty folders (would NOT sync in Linux onedrive):")
	for _, dir := range emptyFolders {
		fmt.Println(dir)
	}
	fmt.Printf("\nTotal empty folders: %d\n", len(emptyFolders))
}
