package env

import "os"

type Config struct {
	Path    ConfigPath
	Scan    ConfigScan
	Tree    ConfigTree
	Folders ConfigFolders
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

type ConfigFolders struct {
	FoldersToScan []os.DirEntry
}

var DefaultConfig = Config{
	Scan: DefaultConfigScan,
	Tree: DefaultConfigTree,
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
