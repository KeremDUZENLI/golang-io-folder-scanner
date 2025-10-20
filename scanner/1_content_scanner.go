package scanner

import (
	"fmt"
	"os"
	"strings"
)

type pathAndContent struct {
	path    string
	content string
}

func ScanFilesContent(files []string) []string {
	pathAndContents := make([]pathAndContent, 0, len(files))

	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		pathAndContents = append(pathAndContents, pathAndContent{file, string(bytes)})
	}

	lines := listPathAndContents(pathAndContents)
	return lines
}

func listPathAndContents(pathAndContents []pathAndContent) []string {
	lines := make([]string, 0, len(pathAndContents)*2)

	for _, pc := range pathAndContents {
		lines = append(lines, fmt.Sprintf("%s=", pc.path))
		lines = append(lines, pc.content)
		lines = append(lines, strings.Repeat("-", 100))
	}

	return lines
}
