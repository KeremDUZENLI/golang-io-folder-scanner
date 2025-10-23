package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/helper"
)

var reader = bufio.NewReader(os.Stdin)

func InputPath(prompt, defaultVal string) string {
	input := readLine(prompt, defaultVal)
	if input == "" {
		return defaultVal
	}

	return helper.CanonicalPath(input)
}

func InputList(prompt string, defaultVal []string) []string {
	input := readLine(prompt, listToString(defaultVal))
	if input == "" {
		return defaultVal
	}
	return lowerStrings(stringToList(input))
}

func InputKeypress() {
	fmt.Print("\nPress ENTER to exit")
	fmt.Scanln()
}

func readLine(prompt, defaultVal string) string {
	fmt.Printf("%s (default = %s): ", prompt, defaultVal)
	input, err := reader.ReadString('\n')
	helper.PrintError(err)
	return strings.TrimSpace(input)
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

func lowerStrings(inputs []string) []string {
	out := make([]string, len(inputs))
	for i, s := range inputs {
		out[i] = strings.ToLower(s)
	}
	return out
}
