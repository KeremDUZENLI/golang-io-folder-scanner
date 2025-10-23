package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

func CanonicalPath(path string) string {
	pathAbs, _ := filepath.Abs(path)
	clean := filepath.Clean(pathAbs)
	return filepath.ToSlash(clean)
}

func PrintError(err error) {
	if err != nil {
		fmt.Printf("\n[ERROR] %v\n", err)
		os.Exit(1)
	}
}
