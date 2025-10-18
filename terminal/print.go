package terminal

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
)

func PrintPathContent(files []scanner.PathContent) {
	printSep()
	fmt.Println("\nSCANNED_FILES=")

	for _, f := range files {
		fmt.Printf("\n%s=\n%s\n", f.Path, f.Content)
		printSep()
	}
}

func PrintTree(lines []string) {
	printSep()
	fmt.Println("\nASCII_TREE=")

	for _, l := range lines {
		fmt.Println(l)
	}
}

func PrintFolders(root string, list []string) {
	printSep()
	fmt.Println("\nFOLDERS=")
	for _, folder := range list {
		pathRel, err := filepath.Rel(root, folder)
		if err != nil {
			pathRel = folder
		}
		fmt.Println(filepath.ToSlash(pathRel))
	}
	fmt.Printf("\nTotal Folders: %d\n", len(list))
}

func PrintError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
	}
}

func printSep() { fmt.Println(strings.Repeat("-", 100)) }
