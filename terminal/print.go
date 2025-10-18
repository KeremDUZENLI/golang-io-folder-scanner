package terminal

import (
	"fmt"
	"path/filepath"
	"strings"
)

func PrintScan(results [][2]string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nSCANNED_FILES=")

	for _, r := range results {
		fmt.Printf("\n%s=\n%s\n", r[0], r[1])
		fmt.Println(strings.Repeat("-", 100))
	}
}

func PrintTree(lines []string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nASCII_TREE=")

	for _, l := range lines {
		fmt.Println(l)
	}
}

func PrintFolders(list []string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nFOLDERS=")
	for _, pathTarg := range list {
		pathRel, err := filepath.Rel(".", pathTarg)
		if err != nil {
			pathRel = pathTarg
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
