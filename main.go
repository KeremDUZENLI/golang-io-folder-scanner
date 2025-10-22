package main

import (
	"fmt"
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func main() {
	cfg := env.ConfigDefault
	cfg.PathToScan, _ = os.Getwd()

	for {
		var choice int
		fmt.Println(env.InputString)
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			cfg.Run_1_FilterFolders()

		case 2:
			cfg.Run_2_FilterFiles()

		case 3:
			cfg.Run_3_ScanFilesContent()

		case 4:
			cfg.Run_4_ScanTree()

		case 5:
			cfg.Run_5_FindFoldersEmpty()

		case 6:
			cfg.Run_6_FindFoldersByFileSuffix()

		case 7:
			cfg.Run_7_CompareFiles()

		default:
			return
		}
	}
}
