package scanner

import "path/filepath"

func GetTreesFromLists(path string, folders, files, foldersContentToSkip []string) []string {
	dirKids := make(map[string][]string)
	fileKids := make(map[string][]string)

	for _, d := range folders {
		if underPath(path, d) {
			parent := filepath.Dir(d)
			dirKids[parent] = append(dirKids[parent], d)
		}
	}

	for _, f := range files {
		if underPath(path, f) {
			parent := filepath.Dir(f)
			fileKids[parent] = append(fileKids[parent], f)
		}
	}

	for k := range dirKids {
		sortPathsByBase(dirKids[k])
	}

	for k := range fileKids {
		sortPathsByBase(fileKids[k])
	}

	lines := []string{filepath.Base(path)}
	return renderTree(path, "", dirKids, fileKids, foldersContentToSkip, lines)
}

func renderTree(parent string, prefix string, dirKids map[string][]string, fileKids map[string][]string, foldersContentToSkip []string, lines []string) []string {
	entries := listEntries(parent, dirKids, fileKids)
	for i, e := range entries {
		lines = append(lines, prefix+treeBranch(i, len(entries))+e.nameBase)
		if e.isDir && !contains(foldersContentToSkip, e.nameBase) {
			nextPrefix := prefix + indent(i, len(entries))
			lines = renderTree(e.pathFull, nextPrefix, dirKids, fileKids, foldersContentToSkip, lines)
		}
	}

	return lines
}
