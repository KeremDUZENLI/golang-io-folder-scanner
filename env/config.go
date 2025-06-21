package env

type Config struct {
	ScanRoot           string
	SuffixesToScan     []string
	SkipFolders        []string
	SkipFoldersContent []string
}

var DefaultConfig = Config{
	SuffixesToScan: []string{
		".py", ".go",
		".html", ".css", ".js",
		".yml", ".json",
	},

	SkipFolders: []string{
		"__pycache__",
		"node_modules",
		"target",
		".git", ".env", ".venv",
		".vscode", ".idea", ".out",
		"dist", "build", "bin", "vendor",
	},

	SkipFoldersContent: []string{
		"data", "img", "images", "assets",
		"logs", "static",
		"tmp", "temp",
	},
}
