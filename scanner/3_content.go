package scanner

import "os"

type FileContent struct {
	Path    string
	Content string
}

func ScanFilesContent(files []string) ([]FileContent, error) {
	results := make([]FileContent, 0, len(files))
	for _, p := range files {
		b, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		results = append(results, FileContent{Path: p, Content: string(b)})
	}
	return results, nil
}
