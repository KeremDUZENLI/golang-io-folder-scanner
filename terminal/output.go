package terminal

import (
	"fmt"
	"strings"
)

func PrintLines(msg string, lines []string) {
	printMsg(msg)

	for _, line := range lines {
		fmt.Println(line)
	}

	printSep()
}

func PrintFolders(msg string, folders []string) {
	printMsg(msg)

	for _, folder := range folders {
		fmt.Println(folder)
	}

	fmt.Printf("\nTOTAL: %d\n", len(folders))
	printSep()
}

func PrintCompare(msg, path1, path2 string, onlyIn1, onlyIn2 []string) {
	printMsg(msg)

	printMsg(fmt.Sprintf("ONLY IN %s", path1))
	for _, p := range onlyIn1 {
		fmt.Println(p)
	}
	fmt.Printf("\nTOTAL: %d\n\n", len(onlyIn1))

	printMsg(fmt.Sprintf("ONLY IN %s", path2))
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
