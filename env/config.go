package env

type Config struct {
	PathToScan        string
	SuffixesToScan    []string
	FoldersToSkip     []string
	FoldersTreeToSkip []string
}

var ConfigDefault = Config{
	PathToScan: ".",

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

	FoldersTreeToSkip: []string{
		"img", "images",
	},
}
