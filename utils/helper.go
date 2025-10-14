package utils

import (
	"bufio"
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

func ReadInput(defaultConfig string) string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	PrintError("Failed to read input", err)

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultConfig
	}
	return input
}

func UpdateListIfInput(input string) []string {
	return strings.Split(strings.ReplaceAll(input, " ", ""), ",")
}

func GetCurrentWorkingDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return cwd, nil
}

func ResolveAbsolutePath(directoryToScan string) (string, error) {
	abs_path, err := filepath.Abs(directoryToScan)
	if err != nil {
		return "", err
	}

	return abs_path, nil
}

func JoinStrings(input []string) string {
	return strings.Join(input, ", ")
}
