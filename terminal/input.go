package terminal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func InputPath(prompt string) string {
	cwd, _ := os.Getwd()
	fmt.Printf("%s (default = %s): ", prompt, cwd)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	PrintError("Failed to read input", err)

	input = strings.TrimSpace(input)
	if input == "" {
		return absolutePathToScan(cwd)
	}

	return absolutePathToScan(input)
}

func absolutePathToScan(directoryToScan string) string {
	absPath, err := filepath.Abs(directoryToScan)
	PrintError("Failed to format path to scan", err)

	return absPath
}

func InputList(prompt string, valueDefault []string) []string {
	fmt.Printf("%s (default = %s): ", prompt, listToString(valueDefault))

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	PrintError("Failed to read input", err)

	input = strings.TrimSpace(input)
	if input == "" {
		return valueDefault
	}

	return stringToList(input)
}

func listToString(list []string) string {
	return strings.Join(list, ", ")
}

func stringToList(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func InputKeypress() {
	fmt.Print("\nPress ENTER to exit")
	fmt.Scanln()
}
