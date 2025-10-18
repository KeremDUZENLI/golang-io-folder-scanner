package env

type Config struct {
	PathToScan        string
	SuffixesToScan    []string
	FoldersToSkip     []string
	FoldersTreeToSkip []string
}

var ConfigDefault = Config{
	PathToScan: ".",
	FoldersToSkip: []string{
		"__pycache__",
		"node_modules",
		"target",
		".git", ".env", ".venv",
		".vscode", ".idea", ".out",
		"dist", "build", "bin", "vendor",
	},

	SuffixesToScan: []string{
		".py", ".go",
		".html", ".css", ".js",
		".yml", ".json",
	},

	FoldersTreeToSkip: []string{
		"img", "images",
	},
}
