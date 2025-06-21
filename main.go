package main

import (
	"fmt"
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
)

func main() {
	cfg := env.DefaultConfig
	scanner.GetUserFilters(&cfg)

	if err := scanner.Traverse(&cfg, cfg.ScanRoot, scanner.HandleFile); err != nil {
		fmt.Fprintf(os.Stderr, "scan error: %v\n", err)
		os.Exit(1)
	}

	if err := scanner.PrintTree(&cfg, cfg.ScanRoot, "", false); err != nil {
		fmt.Fprintf(os.Stderr, "tree error: %v\n", err)
		os.Exit(1)
	}
}
