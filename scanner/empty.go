package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintEmptyFolders(path string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nEMPTY_FOLDERS=")

	emptyFolders := findEmptyFolders(path)

	for _, dir := range emptyFolders {
		if relPath, err := filepath.Rel(".", dir); err == nil {
			fmt.Println(strings.TrimPrefix(relPath, `.\`))
		} else {
			fmt.Println(dir)
		}
	}
	fmt.Printf("\nTotal Empty Folders: %d\n", len(emptyFolders))
}

func findEmptyFolders(path string) []string {
	var emptyFolders []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return emptyFolders
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		pathDir := filepath.Join(path, name)
		if !hasFiles(pathDir) {
			emptyFolders = append(emptyFolders, pathDir)
		}

		emptyFoldersSub := findEmptyFolders(pathDir)
		if len(emptyFoldersSub) > 0 {
			emptyFolders = append(emptyFolders, emptyFoldersSub...)
		}
	}

	return emptyFolders
}
