package env

type Config struct {
	ScanRoot           string
	SuffixesToScan     []string
	SkipFolders        []string
	SkipFoldersContent []string
}

var DefaultConfig = Config{
	SuffixesToScan:     []string{".go"},
	SkipFolders:        []string{"__pycache__", ".venv", ".git", "_scripts"},
	SkipFoldersContent: []string{"data"},
}
