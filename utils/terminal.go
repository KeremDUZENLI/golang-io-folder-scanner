package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadInputPath(prompt, valueDefault string) string {
	fmt.Printf("%s (default = %s): ", prompt, valueDefault)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	PrintError("Failed to Read Input", err)

	input = strings.TrimSpace(input)
	if input == "" {
		return valueDefault
	}

	return formatPathToScan(input)
}

func ReadInputList(prompt string, valueDefault []string) []string {
	fmt.Printf("%s (default = %s): ", prompt, listToString(valueDefault))

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	PrintError("Failed to Read Input", err)

	input = strings.TrimSpace(input)
	if input == "" {
		return valueDefault
	}

	return stringToList(input)
}

func PrintScan(results [][2]string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nSCANNED_FILES=")

	for _, r := range results {
		fmt.Printf("\n%s=\n%s\n", r[0], r[1])
		fmt.Println(strings.Repeat("-", 100))
	}
}

func PrintTree(trees []string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nASCII_TREE=")

	for _, l := range trees {
		fmt.Println(l)
	}
}

func PrintEmptyFolders(emptyFolders []string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nEMPTY_FOLDERS=")

	for _, dir := range emptyFolders {
		normalized := filepath.ToSlash(dir)
		if relPath, err := filepath.Rel(".", dir); err == nil {
			normalized = filepath.ToSlash(relPath)
		}

		fmt.Println(normalized)
	}

	fmt.Printf("\nTotal Empty Folders: %d\n", len(emptyFolders))
}

func PrintError(message string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
		os.Exit(1)
	}
}

func WaitForKeypress() {
	fmt.Print("\nPress ENTER to exit")
	fmt.Scanln()
}
