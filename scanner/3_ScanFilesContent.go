package scanner

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/helper"
)

type Content struct {
	Path    string
	Content string
}

func ScanFilesContent(files []string) []Content {
	contents := make([]Content, 0, len(files))

	for _, file := range files {
		data, err := os.ReadFile(file)
		helper.PrintError(err)

		if filepath.Ext(file) == ".ipynb" {
			data = stripNotebookOutputs(data)
		}

		contents = append(contents, Content{
			Path:    file,
			Content: string(data),
		})
	}

	return contents
}

func stripNotebookOutputs(raw []byte) []byte {
	var nb map[string]interface{}
	json.Unmarshal(raw, &nb)

	cells := nb["cells"].([]interface{})

	for _, c := range cells {
		cell := c.(map[string]interface{})
		delete(cell, "id")
		delete(cell, "metadata")
		delete(cell, "outputs")
		delete(cell, "execution_count")
	}

	out, _ := json.MarshalIndent(nb, "", "  ")
	return out
}
