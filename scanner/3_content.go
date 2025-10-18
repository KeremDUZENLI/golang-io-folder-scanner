package scanner

import "os"

type PathContent struct {
	Path    string
	Content string
}

func ScanFilesPathContent(files []string) ([]PathContent, error) {
	fileContentsList := make([]PathContent, 0, len(files))
	for _, file := range files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		fileContentsList = append(fileContentsList, PathContent{Path: file, Content: string(bytes)})
	}
	return fileContentsList, nil
}
