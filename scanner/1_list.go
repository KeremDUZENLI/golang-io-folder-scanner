package scanner

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ListFolders(root string) ([]string, error) {
	folders := make([]string, 0, 64)
	stack := []string{root}

	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		folders = append(folders, dir)

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		folderNames := make([]string, 0, len(entries))
		for _, entry := range entries {
			if entry.IsDir() {
				folderNames = append(folderNames, entry.Name())
			}
		}

		sortStrings(folderNames)

		for i := len(folderNames) - 1; i >= 0; i-- {
			stack = append(stack, filepath.Join(dir, folderNames[i]))
		}
	}

	return folders, nil
}

func ListFiles(folders []string) ([]string, error) {
	if len(folders) == 0 {
		return nil, nil
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

	return files, nil
}

func listAllowed(folder string, files *[]string, allow map[string]struct{}, visited map[string]struct{}) error {
	if _, ok := allow[folder]; !ok {
		return nil
	}

	key := normDir(folder)
	if _, ok := visited[key]; ok {
		return nil
	}
	visited[key] = struct{}{}

	entries, err := os.ReadDir(folder)
	if err != nil {
		return err
	}

	folders := []string{}
	fileNames := []string{}
	for _, e := range entries {
		if e.IsDir() {
			folders = append(folders, e.Name())
		} else {
			fileNames = append(fileNames, e.Name())
		}
	}

	sortStrings(folders)

	for _, f := range folders {
		_ = listAllowed(filepath.Join(folder, f), files, allow, visited)
	}
	for _, name := range fileNames {
		*files = append(*files, filepath.Join(folder, name))
	}

	return nil
}

func normDir(p string) string {
	return strings.ToLower(filepath.ToSlash(filepath.Clean(p)))
}

func sortStrings(names []string) {
	sort.Slice(names, func(i, j int) bool { return sortAToZ(names[i], names[j]) })
}

func sortAToZ(a, b string) bool {
	return strings.ToLower(a) < strings.ToLower(b)
}
