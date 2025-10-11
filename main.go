package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
)

func main() {
	cfg := env.DefaultConfig
	reader := bufio.NewReader(os.Stdin)

	scanner.GetForScan(reader, &cfg)
	if err := scanner.Traverse(&cfg, cfg.Path.PathToScan, scanner.HandleFile); err != nil {
		fmt.Fprintf(os.Stderr, "scan error: %v\n", err)
		os.Exit(1)
	}

	scanner.GetForTree(reader, &cfg)
	if err := scanner.PrintTree(&cfg, cfg.Path.PathToScan, "", false); err != nil {
		fmt.Fprintf(os.Stderr, "tree error: %v\n", err)
		os.Exit(1)
	}

	scanner.WaitForKeypress()
}
