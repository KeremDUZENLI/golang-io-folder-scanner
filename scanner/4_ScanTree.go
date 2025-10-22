package scanner

import (
	"path"
)

type TreeItem struct {
	Path         string
	IsDir        bool
	Depth        int
	AncestorLast []bool
}

func ScanTree(folders, files, foldersTreeToSkip []string) []TreeItem {
	root := folders[0]
	dirKids := make(map[string][]string, len(folders))
	fileKids := make(map[string][]string, len(files))

	for _, folder := range folders {
		parent := path.Dir(folder)
		dirKids[parent] = append(dirKids[parent], folder)
	}
	for _, file := range files {
		parent := path.Dir(file)
		fileKids[parent] = append(fileKids[parent], file)
	}

	TreeItems := make([]TreeItem, 0, len(folders)+len(files))
	TreeItems = append(TreeItems, TreeItem{
		Path:         root,
		IsDir:        true,
		Depth:        0,
		AncestorLast: nil,
	})

	TreeItems = walkTree(root, 1, nil, dirKids, fileKids, foldersTreeToSkip, TreeItems)
	return TreeItems
}

func walkTree(
	parent string,
	depth int,
	ancestors []bool,
	dirKids,
	fileKids map[string][]string,
	skipNames []string,
	treeItem []TreeItem,
) []TreeItem {
	entries := make([]struct {
		p     string
		isDir bool
	}, 0, len(dirKids[parent])+len(fileKids[parent]))

	for _, d := range dirKids[parent] {
		entries = append(entries, struct {
			p     string
			isDir bool
		}{p: d, isDir: true})
	}
	for _, f := range fileKids[parent] {
		entries = append(entries, struct {
			p     string
			isDir bool
		}{p: f, isDir: false})
	}

	for i, e := range entries {
		isLast := (i == len(entries)-1)
		anc := make([]bool, len(ancestors)+1)
		copy(anc, ancestors)
		anc[len(anc)-1] = isLast

		item := TreeItem{
			Path:         e.p,
			IsDir:        e.isDir,
			Depth:        depth,
			AncestorLast: anc,
		}
		treeItem = append(treeItem, item)

		if e.isDir && !contains(skipNames, path.Base(e.p)) {
			treeItem = walkTree(e.p, depth+1, anc, dirKids, fileKids, skipNames, treeItem)
		}
	}
	return treeItem
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
