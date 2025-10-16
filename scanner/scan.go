package scanner

import (
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func ScanFiles(cfg *env.Config) [][2]string {
	return scanFilesRecursive(cfg.Path.PathToScan, cfg.Path.PathToScan, cfg.Scan.SuffixesToScan, cfg.Scan.FolderToSkip)
}

func scanFilesRecursive(path, pathCurrent string, suffixesToScan, foldersToSkip []string) [][2]string {
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
			results = append(results, scanFilesRecursive(path, pathFull, suffixesToScan, foldersToSkip)...)
			continue
		}

		if hasSuffix(entry.Name(), suffixesToScan) {
			data, err := os.ReadFile(pathFull)
			if err != nil {
				continue
			}

			pathRel, err := filepath.Rel(path, pathFull)
			if err != nil {
				pathRel = pathFull
			}

			results = append(results, [2]string{pathRel, string(data)})
		}
	}

	return results
}
