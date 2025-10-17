package env

type Config struct {
	PathToScan           string
	SuffixesToScan       []string
	FoldersToSkip        []string
	FoldersContentToSkip []string
}

var ConfigDefault = Config{
	SuffixesToScan: []string{
		".py", ".go",
		".html", ".css", ".js",
		".yml", ".json",
	},
	FoldersToSkip: []string{
		"__pycache__",
		"node_modules",
		"target",
		".git", ".env", ".venv",
		".vscode", ".idea", ".out",
		"dist", "build", "bin", "vendor",
	},
	FoldersContentToSkip: []string{
		"img", "images",
	},
}
