package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(prompt, defaultVal string) (string, error) {
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

func WaitForKeypress() {
	fmt.Print("\nPress ENTER to exit")
	fmt.Scanln()
}

func PrintError(message string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
		os.Exit(1)
	}
}
