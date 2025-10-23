package env

import (
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/terminal"
)

func (c *Config) Run_1_FilterFolders() {
	c.PathToScan = terminal.InputPath("[NEW] Path to scan", c.PathToScan)
	folderToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	foldersToSkipAll := append(c.FoldersToSkip, folderToSkip...)

	folders := scanner.ListFolders(c.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)

	terminal.PrintLines("FILTERED FOLDERS", c.PathToScan, foldersFiltered)
}

func (c *Config) Run_2_FilterFiles() {
	c.PathToScan = terminal.InputPath("[NEW] Path to scan", c.PathToScan)
	folderToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)

	foldersToSkipAll := append(c.FoldersToSkip, folderToSkip...)

	folders := scanner.ListFolders(c.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)
	files := scanner.ListFiles(foldersFiltered)
	filesFiltered := scanner.FilterFiles(files, suffixesToScan)

	terminal.PrintLines("FILTERED FILES", c.PathToScan, filesFiltered)
}

func (c *Config) Run_3_ScanFilesContent() {
	c.PathToScan = terminal.InputPath("[NEW] Path to scan", c.PathToScan)
	folderToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)

	foldersToSkipAll := append(c.FoldersToSkip, folderToSkip...)

	folders := scanner.ListFolders(c.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)
	files := scanner.ListFiles(foldersFiltered)
	filesFiltered := scanner.FilterFiles(files, suffixesToScan)
	contents := scanner.ScanFilesContent(filesFiltered)

	terminal.PrintFilesContents("CONTENT OF FILES", c.PathToScan, contents)
}

func (c *Config) Run_4_ScanTree() {
	c.PathToScan = terminal.InputPath("[NEW] Path to scan", c.PathToScan)
	folderToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	foldersTreeToSkip := terminal.InputList("[ADD] Folders tree to skip", c.FoldersTreeToSkip)

	foldersToSkipAll := append(c.FoldersToSkip, folderToSkip...)
	foldersTreeToSkipAll := append(c.FoldersTreeToSkip, foldersTreeToSkip...)

	folders := scanner.ListFolders(c.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)
	files := scanner.ListFiles(foldersFiltered)
	treeItems := scanner.CreateTree(foldersFiltered, files, foldersTreeToSkipAll)

	terminal.PrintTree("ASCII TREE", c.PathToScan, treeItems)
}

func (c *Config) Run_5_FindFoldersEmpty() {
	c.PathToScan = terminal.InputPath("[NEW] Path to scan", c.PathToScan)
	folderToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)

	foldersToSkipAll := append(c.FoldersToSkip, folderToSkip...)

	folders := scanner.ListFolders(c.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)
	foundEmpty := scanner.FindFoldersEmpty(foldersFiltered)

	terminal.PrintLines("EMPTY FOLDERS", c.PathToScan, foundEmpty)
}

func (c *Config) Run_6_FindFoldersByFileSuffix() {
	c.PathToScan = terminal.InputPath("[NEW] Path to scan", c.PathToScan)
	folderToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)

	foldersToSkipAll := append(c.FoldersToSkip, folderToSkip...)

	folders := scanner.ListFolders(c.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)
	foundByFileSuffix := scanner.FindFoldersByFileSuffix(foldersFiltered, suffixesToScan)

	terminal.PrintLines("FOUND FOLDERS", c.PathToScan, foundByFileSuffix)
}

func (c *Config) Run_7_CompareFiles() {
	base1 := terminal.InputPath("1.Path to scan", c.PathToScan)
	base2 := terminal.InputPath("2.Path to scan", c.PathToScan)
	folderToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)

	foldersToSkipAll := append(c.FoldersToSkip, folderToSkip...)

	folders1 := scanner.ListFolders(base1)
	folders2 := scanner.ListFolders(base2)
	foldersFiltered1 := scanner.FilterFolders(folders1, foldersToSkipAll)
	foldersFiltered2 := scanner.FilterFolders(folders2, foldersToSkipAll)
	filesFiltered1 := scanner.ListFiles(foldersFiltered1)
	filesFiltered2 := scanner.ListFiles(foldersFiltered2)
	onlyIn1, onlyIn2 := scanner.CompareFiles(base1, base2, filesFiltered1, filesFiltered2)

	terminal.PrintCompare("FILE COMPARISON", base1, base2, onlyIn1, onlyIn2)
}

func (c *Config) Run_Tester(input string) {
	folders := scanner.ListFolders(c.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, c.FoldersToSkip)
	files := scanner.ListFiles(foldersFiltered)
	filesFiltered := scanner.FilterFiles(files, c.SuffixesToScan)
	inputN := strings.Split(input, ",")

	for _, number := range inputN {
		number = strings.TrimSpace(number)
		switch number {
		case "1":
			for _, i := range foldersFiltered {
				println(i)
			}
		case "2":
			for _, i := range filesFiltered {
				println(i)
			}
		case "3":
			lines := scanner.ScanFilesContent(filesFiltered)
			for _, i := range lines {
				println(i.Path)
				println(i.Content)
			}
		case "4":
			tree := scanner.CreateTree(foldersFiltered, filesFiltered, c.FoldersTreeToSkip)
			for _, i := range tree {
				println(i.Path)
				println(i.IsDir)
				println(i.Depth)
				println(i.AncestorLast)
			}
		case "5":
			foldersEmpty := scanner.FindFoldersEmpty(foldersFiltered)
			for _, i := range foldersEmpty {
				println(i)
			}
		case "6":
			foldersByFileSuffix := scanner.FindFoldersByFileSuffix(foldersFiltered, c.SuffixesToScan)
			for _, i := range foldersByFileSuffix {
				println(i)
			}
		default:
			println("Unknown option:", number)
		}

		println(strings.Repeat("*", 100))
	}
}
