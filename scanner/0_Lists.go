package scanner

import (
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/helper"
)

func ListFolders(path string) []string {
	path = helper.CanonicalPath(path)

	folders := make([]string, 0, 64)
	stack := []string{path}

	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		folders = append(folders, dir)

		fd, err := os.Open(dir)
		helper.PrintError(err)
		entries, err := fd.ReadDir(-1) // DO NOT sort; native order defines traversal & file order
		helper.PrintError(err)
		fd.Close()

		childDirs := make([]string, 0, 8)
		for _, e := range entries {
			if e.IsDir() {
				childDirs = append(childDirs, helper.CanonicalPath(filepath.Join(dir, e.Name())))
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
		listAllowedFiles(folder, &files, allow, visited)
	}

	return files
}

func listAllowedFiles(folder string, files *[]string, allow map[string]struct{}, visited map[string]struct{}) {
	if _, ok := allow[folder]; !ok {
		return
	}
	if _, ok := visited[folder]; ok {
		return
	}
	visited[folder] = struct{}{}

	fd, err := os.Open(folder)
	helper.PrintError(err)

	entries, err := fd.ReadDir(-1) // DO NOT sort; native order defines traversal & file order
	helper.PrintError(err)
	fd.Close()

	dirs := make([]string, 0, 8)
	filesInThisDir := make([]string, 0, 16)

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			dirs = append(dirs, helper.CanonicalPath(filepath.Join(folder, name)))
		} else {
			filesInThisDir = append(filesInThisDir, helper.CanonicalPath(filepath.Join(folder, name)))
		}
	}

	// Recurse into subdirs in left-to-right native order; then emit this dir’s files
	for i := 0; i < len(dirs); i++ {
		listAllowedFiles(dirs[i], files, allow, visited)
	}

	*files = append(*files, filesInThisDir...)
}
