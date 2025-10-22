package scanner

import (
	"path/filepath"
	"strings"
)

func FilterFiles(files, suffixesToScan []string) []string {
	filteredFilesBySuffix := make([]string, 0, len(files))
	for _, file := range files {
		if hasSuffixFile(file, suffixesToScan) {
			filteredFilesBySuffix = append(filteredFilesBySuffix, file)
		}
	}

	return filteredFilesBySuffix
}

func hasSuffixFile(file string, suffixesToScan []string) bool {
	base := filepath.Base(file)
	for _, suffix := range suffixesToScan {
		if strings.HasSuffix(base, suffix) {
			return true
		}
	}

	return false
}
