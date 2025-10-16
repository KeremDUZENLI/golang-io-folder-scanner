package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func formatPathToScan(directoryToScan string) (string, error) {
	return filepath.Abs(directoryToScan)
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

func listToString(list []string) string {
	return strings.Join(list, ", ")
}

func readInput(prompt, defaultVal string) (string, error) {
	fmt.Printf("%s (default = %s): ", prompt, defaultVal)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultVal, nil
	}

	return input, nil
}

func PrintError(message string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
		os.Exit(1)
	}
}
