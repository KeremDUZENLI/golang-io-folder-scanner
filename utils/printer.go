package utils

import (
	"fmt"
	"os"
)

func PrintPrompt(prompt, defaultValues string) {
	fmt.Printf("%s (default = %s): ", prompt, defaultValues)
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
