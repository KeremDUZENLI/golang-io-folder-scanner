package env

type Config struct {
	Path ConfigPath
	Scan ConfigScan
	Tree ConfigTree
}

type ConfigPath struct {
	PathToScan string
}

type ConfigScan struct {
	SuffixesToScan  []string
	FoldersToSkip   []string
	FolderToSkipAdd []string
}

type ConfigTree struct {
	FoldersContentToSkip []string
}

var DefaultConfig = Config{
	Path: DefaultConfigPath,
	Scan: DefaultConfigScan,
	Tree: DefaultConfigTree,
}

var DefaultConfigPath = ConfigPath{
	PathToScan: ".",
}

var DefaultConfigScan = ConfigScan{
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
}

var DefaultConfigTree = ConfigTree{
	FoldersContentToSkip: []string{
		"data", "img", "images", "assets",
		"logs", "static",
		"tmp", "temp",
	},
}
