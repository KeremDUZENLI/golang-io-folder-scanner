package scanner

import (
	"path/filepath"
	"sort"
	"strings"
)

type treeEntry struct {
	nameContent string
	pathFull    string
	isDir       bool
}

func CreateTree(root string, folders, files, foldersTreeToSkip []string) []string {
	dirKids := make(map[string][]string)
	fileKids := make(map[string][]string)

	for _, d := range folders {
		parent := filepath.Dir(d)
		dirKids[parent] = append(dirKids[parent], d)
	}
	for _, f := range files {
		parent := filepath.Dir(f)
		fileKids[parent] = append(fileKids[parent], f)
	}

	for k := range dirKids {
		sortPathsByBase(dirKids[k])
	}
	for k := range fileKids {
		sortPathsByBase(fileKids[k])
	}

	lines := []string{filepath.Base(root)}
	return renderTree(root, "", dirKids, fileKids, foldersTreeToSkip, lines)
}

func sortPathsByBase(paths []string) {
	sort.Slice(paths, func(i, j int) bool {
		ib := strings.ToLower(filepath.Base(paths[i]))
		jb := strings.ToLower(filepath.Base(paths[j]))
		if ib == jb {
			return paths[i] < paths[j]
		}
		return ib < jb
	})
}

func renderTree(parent string, prefix string, dirKids map[string][]string, fileKids map[string][]string, foldersContentToSkip []string, lines []string) []string {
	treeEntryList := listEntries(parent, dirKids, fileKids)
	for i, e := range treeEntryList {
		lines = append(lines, prefix+treeBranch(i, len(treeEntryList))+e.nameContent)
		if e.isDir && !contains(foldersContentToSkip, e.nameContent) {
			nextPrefix := prefix + indent(i, len(treeEntryList))
			lines = renderTree(e.pathFull, nextPrefix, dirKids, fileKids, foldersContentToSkip, lines)
		}
	}
	return lines
}

func listEntries(parent string, dirKids, fileKids map[string][]string) []treeEntry {
	treeEntryList := make([]treeEntry, 0, len(dirKids[parent])+len(fileKids[parent]))
	for _, pathFullDir := range dirKids[parent] {
		treeEntryList = append(treeEntryList, treeEntry{filepath.Base(pathFullDir), pathFullDir, true})
	}
	for _, pathFullFile := range fileKids[parent] {
		treeEntryList = append(treeEntryList, treeEntry{filepath.Base(pathFullFile), pathFullFile, false})
	}
	return treeEntryList
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

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
