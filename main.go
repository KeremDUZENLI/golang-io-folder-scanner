package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	suffixesToScan     = []string{".py"}
	skipFolders        = []string{"__pycache__", ".venv", ".git", "_scripts"}
	skipFoldersContent = []string{"data"}
)

func main() {
	if err := traverse(".", handleFile); err != nil {
		fmt.Fprintf(os.Stderr, "scan error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Project directory structure:")
	if err := printTree(".", "", false); err != nil {
		fmt.Fprintf(os.Stderr, "tree error: %v\n", err)
		os.Exit(1)
	}
}

func handleFile(path string, d fs.DirEntry) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s=\n%s\n", path, data)
	return nil
}

func traverse(root string, handle func(string, fs.DirEntry) error) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		name := d.Name()

		if d.IsDir() {
			if contains(skipFolders, name) {
				return filepath.SkipDir
			}
			if contains(skipFoldersContent, name) {
				return filepath.SkipDir
			}
		}

		if !d.IsDir() && hasAnySuffix(name, suffixesToScan) {
			if contains(skipFoldersContent, filepath.Base(filepath.Dir(path))) {
				return nil
			}
			return handle(path, d)
		}
		return nil
	})
}

func printTree(path, prefix string, skipFiles bool) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	sortEntries(entries)

	var filtered []os.DirEntry
	for _, e := range entries {
		if e.IsDir() && contains(skipFolders, e.Name()) {
			continue
		}
		if skipFiles && !e.IsDir() {
			continue
		}
		filtered = append(filtered, e)
	}

	for i, e := range filtered {
		name := e.Name()
		isDir := e.IsDir()
		fmt.Println(prefix + treeBranch(i, len(filtered)) + name)

		if !isDir {
			continue
		}
		nextSkip := skipFiles || contains(skipFoldersContent, name)
		if err := printTree(
			filepath.Join(path, name),
			prefix+indent(i, len(filtered)),
			nextSkip,
		); err != nil {
			return err
		}
	}
	return nil
}

func sortEntries(es []os.DirEntry) {
	sort.Slice(es, func(i, j int) bool {
		a, b := es[i], es[j]
		if a.IsDir() != b.IsDir() {
			return a.IsDir()
		}
		return strings.ToLower(a.Name()) < strings.ToLower(b.Name())
	})
}

func treeBranch(i, total int) string {
	if i == total-1 {
		return "└── "
	}
	return "├── "
}

func indent(i, total int) string {
	if i == total-1 {
		return "    "
	}
	return "│   "
}

func hasAnySuffix(name string, suffixes []string) bool {
	for _, s := range suffixes {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
