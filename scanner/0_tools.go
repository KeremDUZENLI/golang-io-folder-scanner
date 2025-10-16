package scanner

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func hasFiles(dir string) bool {
	found := false

	walkDir := func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !entry.IsDir() {
			found = true
			return filepath.SkipDir
		}

		return nil
	}

	filepath.WalkDir(dir, walkDir)

	return found
}

func hasSuffix(name string, suffixes []string) bool {
	for _, s := range suffixes {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func sortEntries(entries []os.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		if a.IsDir() == b.IsDir() {
			return strings.ToLower(a.Name()) < strings.ToLower(b.Name())
		}
		return a.IsDir()
	})
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
