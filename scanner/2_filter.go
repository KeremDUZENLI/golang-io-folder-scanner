package scanner

import (
	"path/filepath"
	"strings"
)

func FilterFoldersByName(folders, foldersToSkip []string) []string {
	if len(foldersToSkip) == 0 {
		return folders
	}
	out := make([]string, 0, len(folders))
	for _, f := range folders {
		if skipFolder(f, foldersToSkip) {
			continue
		}
		out = append(out, f)
	}
	return out
}

func skipFolder(path string, foldersToSkip []string) bool {
	if len(foldersToSkip) == 0 {
		return false
	}
	base := strings.ToLower(filepath.Base(path))
	for _, sk := range foldersToSkip {
		if base == strings.ToLower(sk) {
			return true
		}
	}
	return false
}

func FilterFilesBySuffix(foldersFilteredByName, suffixesToScan []string) ([]string, error) {
	files, err := ListFiles(foldersFilteredByName)
	if err != nil {
		return nil, err
	}
	return filterBySuffix(files, suffixesToScan), nil
}

func filterBySuffix(files, suffixesToScan []string) []string {
	if len(suffixesToScan) == 0 {
		return files
	}
	out := make([]string, 0, len(files))
	for _, p := range files {
		if hasSuffix(p, suffixesToScan) {
			out = append(out, p)
		}
	}
	return out
}

func hasSuffix(path string, suffixes []string) bool {
	if len(suffixes) == 0 {
		return true
	}
	base := strings.ToLower(filepath.Base(path))
	for _, s := range suffixes {
		if strings.HasSuffix(base, strings.ToLower(s)) {
			return true
		}
	}
	return false
}
