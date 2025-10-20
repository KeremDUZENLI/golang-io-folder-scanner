package env

type config struct {
	PathToScan        string
	SuffixesToScan    []string
	FoldersToSkip     []string
	FoldersTreeToSkip []string
}

var ConfigDefault = config{
	PathToScan: ".",

	SuffixesToScan: []string{
		".py", ".go",
		".html", ".css", ".js",
		".yml", ".json",
	},

	FoldersToSkip: []string{
		"__pycache__", "node_modules",
		".git", ".gitignore",
		".env", ".venv", ".vscode",
		".idea", ".out",
		"dist", "build", "bin", "vendor", "target",
	},

	FoldersTreeToSkip: []string{
		"img", "images",
	},
}
