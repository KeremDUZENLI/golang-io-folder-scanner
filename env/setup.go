package env

import (
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/terminal"
)

func (c *Config) RunListFolders() []string {
	c.PathToScan = terminal.InputPath("[NEW] Path To scan", c.PathToScan)
	return scanner.ListFolders(c.PathToScan)
}

func (c *Config) RunScanFilesContent(folders []string) {
	foldersToSkip := terminal.InputList("[ADD] Folders to skip", c.FoldersToSkip)
	foldersToSkipAll := append(c.FoldersToSkip, foldersToSkip...)

	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)
	files := scanner.ListFiles(foldersFiltered)

	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)
	filesFiltered := scanner.FilterFiles(files, suffixesToScan)

	lines := scanner.ScanFilesContent(filesFiltered)
	terminal.PrintFilesContents("CONTENT OF FILES", c.PathToScan, lines)
}

func (c *Config) RunTree(folders []string) {
	foldersFiltered := scanner.FilterFolders(folders, c.FoldersToSkip)
	files := scanner.ListFiles(foldersFiltered)

	foldersTreeToSkip := terminal.InputList("[ADD] Folders tree to skip", c.FoldersTreeToSkip)
	foldersTreeToSkipAll := append(c.FoldersTreeToSkip, foldersTreeToSkip...)

	lines := scanner.CreateTree(foldersFiltered, files, foldersTreeToSkipAll)
	terminal.PrintTreeASCII("ASCII TREE", c.PathToScan, lines)
}

func (c *Config) RunFoldersEmpty(folders []string) {
	foldersFiltered := scanner.FilterFolders(folders, c.FoldersToSkip)
	foundEmpty := scanner.FindFoldersEmpty(foldersFiltered)

	terminal.PrintLines("EMPTY FOLDERS", c.PathToScan, foundEmpty)
}

func (c *Config) RunFoldersBySuffix(folders []string) {
	suffixesToScan := terminal.InputList("[NEW] Suffixes to scan", c.SuffixesToScan)
	foldersFiltered := scanner.FilterFolders(folders, c.FoldersToSkip)
	foundByFileSuffix := scanner.FindFoldersByFileSuffix(foldersFiltered, suffixesToScan)

	terminal.PrintLines("FOUND FOLDERS", c.PathToScan, foundByFileSuffix)
}

func (c *Config) RunFilesCompare(folders1, folders2 []string) {
	foldersFiltered1 := scanner.FilterFolders(folders1, c.FoldersToSkip)
	foldersFiltered2 := scanner.FilterFolders(folders2, c.FoldersToSkip)

	filesFiltered1 := scanner.ListFiles(foldersFiltered1)
	filesFiltered2 := scanner.ListFiles(foldersFiltered2)
	onlyIn1, onlyIn2 := scanner.CompareFiles(filesFiltered1, filesFiltered2)

	terminal.PrintCompare("FILE COMPARISON", folders1[0], folders2[0], onlyIn1, onlyIn2)
}

func (c *Config) RunTester(input string) {
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
