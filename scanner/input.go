package scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func ReadInput(reader *bufio.Reader, prompt, defaultConfig string) string {
	fmt.Printf("%s (default = %s): ", prompt, defaultConfig)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetForPath(reader *bufio.Reader) string {
	cwd, _ := os.Getwd()
	dir := ReadInput(reader, "Folder Directory", cwd)

	abs_path, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve: %v\n", err)
		os.Exit(1)
	}

	return abs_path
}

func GetForScan(reader *bufio.Reader, cfg *env.Config) {
	cfg.Path.PathToScan = GetForPath(reader)

	suffixesToScan := ReadInput(reader, "Suffixes To Scan", strings.Join(cfg.Scan.SuffixesToScan, ", "))
	if suffixesToScan != "" {
		cfg.Scan.SuffixesToScan = strings.Split(strings.ReplaceAll(suffixesToScan, " ", ""), ",")
	}

	foldersToSkip := ReadInput(reader, "Skip Folders", strings.Join(cfg.Scan.FoldersToSkip, ", "))
	if foldersToSkip != "" {
		cfg.Scan.FoldersToSkip = strings.Split(strings.ReplaceAll(foldersToSkip, " ", ""), ",")
	}
}

func GetForTree(reader *bufio.Reader, cfg *env.Config) {
	cfg.Path.PathToScan = GetForPath(reader)

	folderContentToSkip := ReadInput(reader, "Skip Folders Content", strings.Join(cfg.Tree.FoldersContentToSkip, ", "))
	if folderContentToSkip != "" {
		cfg.Tree.FoldersContentToSkip = strings.Split(strings.ReplaceAll(folderContentToSkip, " ", ""), ",")
	}
}

func WaitForKeypress() {
	fmt.Print("\nPress ENTER to exit")
	fmt.Scanln()
}
