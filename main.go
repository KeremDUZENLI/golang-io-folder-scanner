package main

import (
	"fmt"
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
)

func main() {
	cfg := env.DefaultConfig
	dir := scanner.GetScanDirectory()
	cfg.ScanRoot = dir

	if err := scanner.Traverse(&cfg, dir, scanner.HandleFile); err != nil {
		fmt.Fprintf(os.Stderr, "scan error: %v\n", err)
		os.Exit(1)
	}

	if err := scanner.PrintTree(&cfg, dir, "", false); err != nil {
		fmt.Fprintf(os.Stderr, "tree error: %v\n", err)
		os.Exit(1)
	}
}
