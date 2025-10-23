package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

func CanonicalPath(base string) string {
	pathAbs, _ := filepath.Abs(base)
	clean := filepath.Clean(pathAbs)
	return filepath.ToSlash(clean)
}

func RelativePath(base, path string) string {
	rel, err := filepath.Rel(base, path)
	if err != nil {
		return path
	}
	return filepath.ToSlash(rel)
}

func PrintError(err error) {
	if err != nil {
		fmt.Printf("\n[ERROR] %v\n", err)
		os.Exit(1)
	}
}
