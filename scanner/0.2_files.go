package scanner

import (
	"os"
	"path/filepath"
)

func ListFiles(folders []string) []string {
	if len(folders) == 0 {
		return nil
	}

	allow := make(map[string]struct{}, len(folders))
	for _, d := range folders {
		allow[d] = struct{}{}
	}

	files := make([]string, 0, 256)
	visited := make(map[string]struct{}, len(folders))

	for _, folder := range folders {
		_ = listAllowed(folder, &files, allow, visited)
	}

	return files
}

func listAllowed(folder string, files *[]string, allow map[string]struct{}, visited map[string]struct{}) error {
	if _, ok := allow[folder]; !ok {
		return nil
	}
	if _, ok := visited[folder]; ok {
		return nil
	}
	visited[folder] = struct{}{}

	fd, err := os.Open(folder)
	if err != nil {
		return err
	}
	entries, err := fd.ReadDir(-1)
	_ = fd.Close()
	if err != nil {
		return err
	}

	dirs := make([]string, 0, 8)
	filesInThisDir := make([]string, 0, 16)

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			dirs = append(dirs, canonicalPath(filepath.Join(folder, name)))
		} else {
			filesInThisDir = append(filesInThisDir, canonicalPath(filepath.Join(folder, name)))
		}
	}
	for i := 0; i < len(dirs); i++ {
		_ = listAllowed(dirs[i], files, allow, visited)
	}

	*files = append(*files, filesInThisDir...)

	return nil
}
