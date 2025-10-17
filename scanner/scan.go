package scanner

import (
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func ScanFiles(cfg *env.Config, cfgAdd *env.ConfigAdd) [][2]string {
	foldersToSkipTotal := append(cfg.FoldersToSkip, cfgAdd.FoldersToSkip...)
	return scanFilesRecursive(cfg.PathToScan, cfg.PathToScan, cfg.SuffixesToScan, foldersToSkipTotal)
}

func scanFilesRecursive(pathToScan, pathCurrent string, suffixesToScan, foldersToSkip []string) [][2]string {
	var results [][2]string

	entries, err := os.ReadDir(pathCurrent)
	if err != nil {
		return results
	}

	sortEntries(entries)

	for _, entry := range entries {
		pathFull := filepath.Join(pathCurrent, entry.Name())

		if entry.IsDir() {
			if contain(foldersToSkip, entry.Name()) {
				continue
			}
			results = append(results, scanFilesRecursive(pathToScan, pathFull, suffixesToScan, foldersToSkip)...)
			continue
		}

		if hasSuffix(entry.Name(), suffixesToScan) {
			data, err := os.ReadFile(pathFull)
			if err != nil {
				continue
			}

			pathRel, err := filepath.Rel(pathToScan, pathFull)
			if err != nil {
				pathRel = pathFull
			}

			results = append(results, [2]string{pathRel, string(data)})
		}
	}

	return results
}
