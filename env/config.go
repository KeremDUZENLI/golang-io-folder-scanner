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

var InputString string = `
Welcome to io-folder-scanner

1) List folders
2) List files
3) Scan content of files
4) Create ASCII tree
5) Find folders empty
6) Find folders by file suffix
7) Compare two paths

Press ENTER to quit

Choose [1-7]:`
