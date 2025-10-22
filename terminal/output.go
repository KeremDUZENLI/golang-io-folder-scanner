package terminal

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
)

func PrintLines(msg, base string, lines []string) {
	printMsg(msg)

	for _, line := range lines {
		fmt.Println(relativePath(base, line))
	}

	fmt.Printf("\nTOTAL: %d\n", len(lines))
	printSep()
}

func PrintFilesContents(msg, base string, lines []scanner.Content) {
	printMsg(msg)

	for _, line := range lines {
		fmt.Println(relativePath(base, line.Path) + "=")
		fmt.Println(line.Content)
		fmt.Println(strings.Repeat("-", 100))
	}

	printSep()
}

func PrintTree(msg, base string, lines []scanner.TreeItem) {
	printMsg(msg)

	for _, line := range lines[1:] {
		prefix := buildGuides(line.AncestorLast)
		branch := "├── "
		if len(line.AncestorLast) > 0 && line.AncestorLast[len(line.AncestorLast)-1] {
			branch = "└── "
		}
		path := filepath.Base(relativePath(base, line.Path))
		fmt.Println(prefix + branch + path)
	}

	printSep()
}

func PrintCompare(msg, path1, path2 string, onlyIn1, onlyIn2 []string) {
	printMsg(msg)

	printMsg(fmt.Sprintf("ONLY IN %s", path1))
	for _, p := range onlyIn1 {
		fmt.Println(relativePath(path1, p))
	}
	fmt.Printf("\nTOTAL: %d\n\n", len(onlyIn1))

	printMsg(fmt.Sprintf("ONLY IN %s", path2))
	for _, p := range onlyIn2 {
		fmt.Println(relativePath(path2, p))
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

func relativePath(base, path string) string {
	rel, err := filepath.Rel(base, path)
	if err != nil {
		return path
	}
	return filepath.ToSlash(rel)
}

func buildGuides(ancestorLast []bool) string {
	var b strings.Builder
	for i := 0; i < len(ancestorLast)-1; i++ {
		if ancestorLast[i] {
			b.WriteString("    ")
		} else {
			b.WriteString("│   ")
		}
	}
	return b.String()
}
