package main

import (
	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/utils"
)

func main() {
	cfg := env.DefaultConfig

	utils.GetForPath(&cfg)

	utils.GetForScan(&cfg)
	scanner.PrintScan(&cfg, cfg.Path.PathToScan)

	utils.GetForTree(&cfg)
	scanner.PrintTree(&cfg, cfg.Path.PathToScan)

	utils.WaitForKeypress()
}
