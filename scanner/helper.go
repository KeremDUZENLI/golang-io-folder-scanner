package scanner

import (
	"path/filepath"
	"sort"
	"strings"
)

type entry struct {
	nameBase string
	pathFull string
	isDir    bool
}

func listEntries(parent string, dirKids, fileKids map[string][]string) []entry {
	entries := make([]entry, 0, len(dirKids[parent])+len(fileKids[parent]))

	for _, d := range dirKids[parent] {
		entries = append(entries, entry{
			nameBase: filepath.Base(d),
			pathFull: d,
			isDir:    true,
		})
	}
	for _, f := range fileKids[parent] {
		entries = append(entries, entry{
			nameBase: filepath.Base(f),
			pathFull: f,
			isDir:    false,
		})
	}
	return entries
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func underPath(root, child string) bool {
	rel, err := filepath.Rel(root, child)
	return err == nil && !strings.HasPrefix(rel, "..")
}

func sortPathsByBase(paths []string) {
	sort.Slice(paths, func(i, j int) bool {
		return strings.ToLower(filepath.Base(paths[i])) < strings.ToLower(filepath.Base(paths[j]))
	})
}

func hasSuffix(path string, suffixes []string) bool {
	if len(suffixes) == 0 {
		return true
	}

	pathBase := strings.ToLower(filepath.Base(path))
	for _, s := range suffixes {
		if strings.HasSuffix(pathBase, strings.ToLower(s)) {
			return true
		}
	}

	return false
}

func skipFolder(path string, foldersToSkip []string) bool {
	if len(foldersToSkip) == 0 {
		return false
	}

	pathBase := strings.ToLower(filepath.Base(path))
	for _, folderToSkip := range foldersToSkip {
		if pathBase == strings.ToLower(folderToSkip) {
			return true
		}
	}

	return false
}

func treeBranch(index, total int) string {
	if index == total-1 {
		return "└── "
	}
	return "├── "
}

func indent(index, total int) string {
	if index == total-1 {
		return "    "
	}
	return "│   "
}
