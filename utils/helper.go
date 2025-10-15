package utils

import (
	"os"
	"sort"
	"strings"
)

func Contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func HasAnySuffix(name string, suffixes []string) bool {
	for _, s := range suffixes {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

func SortEntries(entries []os.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		if a.IsDir() == b.IsDir() {
			return strings.ToLower(a.Name()) < strings.ToLower(b.Name())
		}
		return a.IsDir()
	})
}

func TreeBranch(index, total int) string {
	if index == total-1 {
		return "└── "
	}
	return "├── "
}

func Indent(index, total int) string {
	if index == total-1 {
		return "    "
	}
	return "│   "
}
