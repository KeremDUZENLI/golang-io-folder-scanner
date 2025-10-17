package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

type FileResult struct {
	Path    string
	Content string
}

func ListFolders(root string, foldersToSkip []string) ([]string, error) {
	var folders []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}
		if !d.IsDir() {
			return nil
		}

		if path != root {
			name := d.Name()
			if contains(foldersToSkip, name) {
				return filepath.SkipDir
			}
		}

		folders = append(folders, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func ListFiles(root string, suffixesToScan, foldersToSkip []string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}
		if d.IsDir() {
			if path != root {
				if contains(foldersToSkip, d.Name()) {
					return filepath.SkipDir
				}
			}
			return nil
		}

		if len(suffixesToScan) == 0 || hasAnySuffix(strings.ToLower(d.Name()), suffixesToScan) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ReadFiles(paths []string) ([]FileResult, error) {
	out := make([]FileResult, 0, len(paths))
	for _, p := range paths {
		b, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		out = append(out, FileResult{Path: p, Content: string(b)})
	}
	return out, nil
}
