package scanner

import (
	"path/filepath"
)

type treeEntry struct {
	pathFull    string
	contentName string
	isDir       bool
}

func CreateTree(folders, files, foldersTreeToSkip []string) []string {
	root := folders[0]
	dirKids := make(map[string][]string, len(folders))
	fileKids := make(map[string][]string, len(files))

	for _, d := range folders {
		parent := filepath.Dir(d)
		dirKids[parent] = append(dirKids[parent], d)
	}

	for _, f := range files {
		parent := filepath.Dir(f)
		fileKids[parent] = append(fileKids[parent], f)
	}

	lines := []string{filepath.Base(root)}
	return renderTree(root, "", dirKids, fileKids, foldersTreeToSkip, lines)
}

func renderTree(parent, prefix string, dirKids, fileKids map[string][]string, foldersContentToSkip, lines []string) []string {
	entries := listTreeEntries(parent, dirKids, fileKids)

	for i, e := range entries {
		lines = append(lines, prefix+treeBranch(i, len(entries))+e.contentName)
		if e.isDir && !contains(foldersContentToSkip, e.contentName) {
			nextPrefix := prefix + indent(i, len(entries))
			lines = renderTree(e.pathFull, nextPrefix, dirKids, fileKids, foldersContentToSkip, lines)
		}
	}

	return lines
}

func listTreeEntries(parent string, dirKids, fileKids map[string][]string) []treeEntry {
	entries := make([]treeEntry, 0, len(dirKids[parent])+len(fileKids[parent]))

	for _, d := range dirKids[parent] {
		entries = append(entries, treeEntry{
			pathFull:    d,
			contentName: filepath.Base(d),
			isDir:       true,
		})
	}

	for _, f := range fileKids[parent] {
		entries = append(entries, treeEntry{
			pathFull:    f,
			contentName: filepath.Base(f),
			isDir:       false,
		})
	}

	return entries
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}

	return false
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
