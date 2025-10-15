package utils

import (
	"os"
	"path/filepath"
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

func StringToList(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func ListToString(list []string) string {
	return strings.Join(list, ", ")
}

func GetCurrentWorkingDirectory() (string, error) {
	return os.Getwd()
}

func FormatPathToScan(directoryToScan string) (string, error) {
	return filepath.Abs(directoryToScan)
}
