package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getCurrentWorkingDirectory() (string, error) {
	return os.Getwd()
}

func formatPathToScan(directoryToScan string) (string, error) {
	return filepath.Abs(directoryToScan)
}

func getFoldersToScan(directoryToScan string) ([]os.DirEntry, error) {
	return os.ReadDir(directoryToScan)
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
