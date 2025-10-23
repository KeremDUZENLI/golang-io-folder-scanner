package scanner

import (
	"os"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/helper"
)

type Content struct {
	Path    string
	Content string
}

func ScanFilesContent(files []string) []Content {
	contents := make([]Content, 0, len(files))

	for _, file := range files {
		bytes, err := os.ReadFile(file)
		helper.PrintError(err)
		contents = append(contents, Content{
			Path:    file,
			Content: string(bytes),
		})
	}
	return contents
}
