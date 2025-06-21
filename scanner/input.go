package scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetScanDirectory() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Folder Directory: ")
	input, _ := reader.ReadString('\n')
	dir := strings.TrimSpace(input)

	absPath, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve path: %v\n", err)
		os.Exit(1)
	}

	stat, err := os.Stat(absPath)
	if err != nil || !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "Invalid directory: %v\n", absPath)
		os.Exit(1)
	}

	return absPath
}
