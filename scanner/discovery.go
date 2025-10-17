package scanner

import (
	"io/fs"
	"os"
	"path/filepath"
)

type FileResult struct {
	Path    string
	Content string
}

func ReadFiles(paths []string) ([]FileResult, error) {
	contents := make([]FileResult, 0, len(paths))
	for _, path := range paths {
		b, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		contents = append(contents, FileResult{Path: path, Content: string(b)})
	}
	return contents, nil
}

func ScanFoldersAndFiles(path string) (folders []string, files []string, err error) {
	err = filepath.WalkDir(path, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if d.IsDir() {
			folders = append(folders, p)
		} else {
			files = append(files, p)
		}
		return nil
	})

	return folders, files, err
}
