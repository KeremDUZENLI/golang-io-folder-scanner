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
	SuffixesToScan       []string
	DefaultFoldersToSkip []string
	FolderToSkip         []string
}

type ConfigTree struct {
	DefaultFoldersContentToSkip []string
	FoldersContentToSkip        []string
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

	DefaultFoldersToSkip: []string{
		"__pycache__",
		"node_modules",
		"target",
		".git", ".env", ".venv",
		".vscode", ".idea", ".out",
		"dist", "build", "bin", "vendor",
	},
}

var DefaultConfigTree = ConfigTree{
	DefaultFoldersContentToSkip: []string{
		"img", "images",
	},
}
