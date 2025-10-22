package helper

import "path/filepath"

func CanonicalPath(path string) string {
	pathAbs, _ := filepath.Abs(path)
	clean := filepath.Clean(pathAbs)
	return filepath.ToSlash(clean)
}
