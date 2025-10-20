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

	files := []string{}
	root := folders[0]

	err := listFilesRecursive(root, &files)
	return files, err
}

func listFilesRecursive(dir string, files *[]string) error {
	entries, err := os.ReadDir(dir)
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
	sortStrings(fileNames)

	for _, folder := range folders {
		subdir := filepath.Join(dir, folder)
		if err := listFilesRecursive(subdir, files); err != nil {
			continue
		}
	}
	for _, name := range fileNames {
		*files = append(*files, filepath.Join(dir, name))
	}

	return nil
}

func sortStrings(names []string) {
	sort.Slice(names, func(i, j int) bool { return sortAToZ(names[i], names[j]) })
}

func sortAToZ(a, b string) bool {
	return strings.ToLower(a) < strings.ToLower(b)
}
