package terminal

import (
	"fmt"
	"path/filepath"
	"strings"
)

func PrintLines(msg string, lines []string) {
	printMsg(msg)

	for _, line := range lines {
		fmt.Println(line)
	}

	printSep()
}

func PrintFolders(msg, root string, lines []string) {
	printMsg(msg)

	for _, folder := range lines {
		pathRel, err := filepath.Rel(root, folder)
		if err != nil {
			pathRel = folder
		}
		fmt.Println(filepath.ToSlash(pathRel))
	}

	fmt.Printf("\nTOTAL: %d\n", len(lines))
	printSep()
}

func PrintCompare(msg, path1, path2 string, onlyIn1, onlyIn2 []string) {
	printMsg(msg)
	fmt.Printf("PATH 1: %s\nPATH 2: %s\n\n", path1, path2)

	fmt.Println("[ONLY IN PATH 1]")
	for _, p := range onlyIn1 {
		fmt.Println(p)
	}
	fmt.Printf("\nTOTAL: %d\n\n", len(onlyIn1))

	fmt.Println("[ONLY IN PATH 2]")
	for _, p := range onlyIn2 {
		fmt.Println(p)
	}
	fmt.Printf("\nTOTAL: %d\n", len(onlyIn2))

	printSep()
}

func PrintError(msg string, err error) {
	if err != nil {
		fmt.Printf("\n%s: %v\n", msg, err)
	}
}

func printMsg(msg string) {
	left := (100 - len(msg)) / 2
	right := 100 - len(msg) - left
	fmt.Printf("\n%s%s%s\n\n", strings.Repeat("_", left), msg, strings.Repeat("_", right))
}

func printSep() {
	fmt.Printf("\n%s\n\n", strings.Repeat("_", 100))
}
