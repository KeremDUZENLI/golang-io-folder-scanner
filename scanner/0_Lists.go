package scanner

import (
	"os"
	"path/filepath"
)

func ListFolders(path string) []string {
	path = canonicalPath(path)

	folders := make([]string, 0, 64)
	stack := []string{path}

	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		folders = append(folders, dir)

		fd, err := os.Open(dir)
		if err != nil {
			continue
		}
		defer fd.Close()

		entries, err := fd.ReadDir(-1) // DO NOT sort; native order defines traversal & file order
		if err != nil {
			continue
		}

		childDirs := make([]string, 0, 8)
		for _, e := range entries {
			if e.IsDir() {
				childDirs = append(childDirs, canonicalPath(filepath.Join(dir, e.Name())))
			}
		}
		for i := len(childDirs) - 1; i >= 0; i-- {
			stack = append(stack, childDirs[i])
		}
	}

	return folders
}

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
		_ = listAllowedFiles(folder, &files, allow, visited)
	}

	return files
}

func listAllowedFiles(folder string, files *[]string, allow map[string]struct{}, visited map[string]struct{}) error {
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
	defer fd.Close()

	entries, err := fd.ReadDir(-1) // DO NOT sort; native order defines traversal & file order
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

	// Recurse into subdirs in left-to-right native order; then emit this dir’s files
	for i := 0; i < len(dirs); i++ {
		_ = listAllowedFiles(dirs[i], files, allow, visited)
	}

	*files = append(*files, filesInThisDir...)

	return nil
}

func canonicalPath(path string) string {
	pathAbs, _ := filepath.Abs(path)
	clean := filepath.Clean(pathAbs)
	return filepath.ToSlash(clean)
}
