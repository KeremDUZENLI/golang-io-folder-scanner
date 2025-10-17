package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadInputPath(prompt string) string {
	cwd, _ := os.Getwd()
	fmt.Printf("%s (default = %s): ", prompt, cwd)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	PrintError("Failed to Read Input", err)

	input = strings.TrimSpace(input)
	if input == "" {
		return absolutePathToScan(cwd)
	}

	return absolutePathToScan(input)
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

func PrintTree(lines []string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nASCII_TREE=")

	for _, l := range lines {
		fmt.Println(l)
	}
}

func PrintEmptyFolders(list []string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("\nEMPTY_FOLDERS=")
	for _, pathTarg := range list {
		pathRel, err := filepath.Rel(".", pathTarg)
		if err != nil {
			pathRel = pathTarg
		}
		fmt.Println(filepath.ToSlash(pathRel))
	}

	fmt.Printf("\nTotal Empty Folders: %d\n", len(list))
}

func PrintError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
	}
}

func WaitForKeypress() {
	fmt.Print("\nPress ENTER to exit")
	fmt.Scanln()
}
