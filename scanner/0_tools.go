package scanner

import (
	"os"
	"sort"
	"strings"
)

func hasSuffix(name string, suffixes []string) bool {
	for _, s := range suffixes {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

func contain(list []string, s string) bool {
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
