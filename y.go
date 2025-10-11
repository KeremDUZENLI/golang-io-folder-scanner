package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func hasFiles(dir string) bool {
	hasFile := false
	err := filepath.WalkDir(dir, func(path string, entry os.DirEntry, err error) error {

		if !entry.IsDir() {
			hasFile = true
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return true
	}

	return hasFile
}

func findEmptyDirs(root string) []string {
	var emptyDirs []string

	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading root folder:", err)
		return emptyDirs
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirPath := filepath.Join(root, entry.Name())

			if !hasFiles(dirPath) {
				emptyDirs = append(emptyDirs, dirPath)
			}

			subEmpty := findEmptyDirs(dirPath)
			emptyDirs = append(emptyDirs, subEmpty...)
		}
	}

	return emptyDirs
}

func printEmptyDirs(emptyDirs []string) {
	if len(emptyDirs) == 0 {
		fmt.Println("No empty folders found.")
		return
	}

	fmt.Println("Empty folders (would NOT sync in Linux onedrive):")
	for _, dir := range emptyDirs {
		fmt.Println(dir)
	}
	fmt.Printf("\nTotal empty folders: %d\n", len(emptyDirs))
}

func main() {
	root := `D:\OneDrive - Debreceni Egyetem`
	emptyDirs := findEmptyDirs(root)
	printEmptyDirs(emptyDirs)
}
