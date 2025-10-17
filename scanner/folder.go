package scanner

import (
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func FindEmptyFolders(cfg *env.Config) []string {
	return findEmptyFoldersRecursive(cfg.PathToScan, cfg.PathToScan, cfg.FoldersToSkip)
}

func findEmptyFoldersRecursive(path, pathCurrent string, defaultFoldersToSkip []string) []string {
	var emptyFolders []string

	entries, err := os.ReadDir(pathCurrent)
	if err != nil {
		return emptyFolders
	}

	hasFile := false
	for _, e := range entries {
		if e.IsDir() {
			if contain(defaultFoldersToSkip, e.Name()) {
				continue
			}

			pathSub := filepath.Join(pathCurrent, e.Name())
			emptyFoldersSub := findEmptyFoldersRecursive(path, pathSub, defaultFoldersToSkip)
			emptyFolders = append(emptyFolders, emptyFoldersSub...)

			if !contain(emptyFolders, filepath.ToSlash(pathSub)) {
				hasFile = true
			}
		} else {
			hasFile = true
		}
	}

	if !hasFile {
		relPath, err := filepath.Rel(path, pathCurrent)
		if err != nil {
			relPath = pathCurrent
		}
		emptyFolders = append(emptyFolders, filepath.ToSlash(relPath))
	}

	return emptyFolders
}
