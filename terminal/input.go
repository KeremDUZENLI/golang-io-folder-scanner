package terminal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func InputPath(prompt string, defaultVal string) string {
	input := readLine(prompt, defaultVal)
	if input == "" {
		return absolutePathToScan(defaultVal)
	}
	return absolutePathToScan(input)
}

func InputList(prompt string, defaultVal []string) []string {
	input := readLine(prompt, listToString(defaultVal))
	if input == "" {
		return defaultVal
	}
	return stringToList(input)
}

func InputKeypress() {
	fmt.Print("\nPress ENTER to exit")
	fmt.Scanln()
}

func readLine(prompt, defaultVal string) string {
	fmt.Printf("%s (default = %s): ", prompt, defaultVal)
	input, err := reader.ReadString('\n')
	PrintError("Failed to read input", err)
	return strings.TrimSpace(input)
}

func absolutePathToScan(directoryToScan string) string {
	absPath, err := filepath.Abs(directoryToScan)
	PrintError("Failed to format path to scan", err)

	return absPath
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
