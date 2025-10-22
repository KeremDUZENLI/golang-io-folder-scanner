package scanner

import (
	"os"
)

type Content struct {
	Path    string
	Content string
}

func ScanFilesContent(files []string) []Content {
	contents := make([]Content, 0, len(files))

	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		contents = append(contents, Content{
			Path:    file,
			Content: string(bytes),
		})
	}
	return contents
}
