package main

import (
	"fmt"
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
)

func main() {
	if err := scanner.Traverse(".", scanner.HandleFile); err != nil {
		fmt.Fprintf(os.Stderr, "scan error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Project directory structure:")
	if err := scanner.PrintTree(".", "", false); err != nil {
		fmt.Fprintf(os.Stderr, "tree error: %v\n", err)
		os.Exit(1)
	}
}
