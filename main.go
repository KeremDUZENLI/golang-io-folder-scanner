package main

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func main() {
	cfg := env.DefaultConfig

	scanner.GetForPath(&cfg)
	scanner.GetForScan(&cfg)
	scanner.Traverse(&cfg, cfg.Path.PathToScan, scanner.HandleFile)
	scanner.PrintTree(&cfg, cfg.Path.PathToScan, "", false)

	utils.WaitForKeypress()
}
