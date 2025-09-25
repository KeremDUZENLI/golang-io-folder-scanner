package scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
)

func GetUserFilters(cfg *env.Config) {
	reader := bufio.NewReader(os.Stdin)
	cwd, _ := os.Getwd()

	fmt.Printf("Folder Directory: (default = %s): ", cwd)
	dirLine, _ := reader.ReadString('\n')
	dir := strings.TrimSpace(dirLine)
	abs_path, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve: %v\n", err)
		os.Exit(1)
	}
	if stat, err := os.Stat(abs_path); err != nil || !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "Invalid directory: %v\n", abs_path)
		os.Exit(1)
	}
	cfg.ScanRoot = abs_path

	fmt.Printf("Suffixes To Scan (default = %s): ", strings.Join(cfg.SuffixesToScan, ", "))
	line, _ := reader.ReadString('\n')
	if s := strings.TrimSpace(line); s != "" {
		parts := strings.Split(s, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		cfg.SuffixesToScan = parts
	}

	fmt.Printf("Skip Folders (default = %s): ", strings.Join(cfg.SkipFolders, ", "))
	line, _ = reader.ReadString('\n')
	if s := strings.TrimSpace(line); s != "" {
		parts := strings.Split(s, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		cfg.SkipFolders = parts
	}

	fmt.Printf("Skip Folders Content (default = %s): ", strings.Join(cfg.SkipFoldersContent, ", "))
	line, _ = reader.ReadString('\n')
	if s := strings.TrimSpace(line); s != "" {
		parts := strings.Split(s, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		cfg.SkipFoldersContent = parts
	}
}

func WaitForKeypress() {
	fmt.Print("\nPress ENTER to exit")
	fmt.Scanln()
}
