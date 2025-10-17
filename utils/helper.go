package utils

import (
	"path/filepath"
	"strings"
)

func formatPathToScan(directoryToScan string) string {
	absPath, err := filepath.Abs(directoryToScan)
	PrintError("Failed to Format Path to Scan", err)

	return absPath
}

func listToString(list []string) string {
	return strings.Join(list, ", ")
}

func stringToList(s string) []string {
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
