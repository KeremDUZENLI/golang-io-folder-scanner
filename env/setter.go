package env

func (c *Config) SetForPath() {
	cwd, err := getCurrentWorkingDirectory()
	PrintError("Failed to Get Current Working Directory", err)

	input, err := readInput("[OVERRIDE] Path To Scan", cwd)
	PrintError("Failed to Read Path To Scan", err)

	pathToScan, err := formatPathToScan(input)
	PrintError("Failed to Format Path to Scan", err)

	c.Path.PathToScan = pathToScan
}

func (c *Config) SetForScan() {
	defaultSuffixesToScan := listToString(c.Scan.SuffixesToScan)
	input, err := readInput("[OVERRIDE] Suffixes to Scan", defaultSuffixesToScan)
	PrintError("Failed to Read Suffixes to Scan", err)
	c.Scan.SuffixesToScan = stringToList(input)

	defaultFoldersToSkip := listToString(c.Scan.DefaultFoldersToSkip)
	input, err = readInput("[ADD] Folders to Skip", defaultFoldersToSkip)
	PrintError("Failed to Read Folders to Skip", err)
	c.Scan.FolderToSkip = stringToList(defaultFoldersToSkip + "," + input)
}

func (c *Config) SetForTree() {
	defaultFoldersContentToSkip := listToString(c.Tree.DefaultFoldersContentToSkip)
	input, err := readInput("[ADD] Folders Content to Skip", defaultFoldersContentToSkip)
	PrintError("Failed to Read Folders Content to Skip", err)
	c.Tree.FoldersContentToSkip = stringToList(defaultFoldersContentToSkip + "," + input)
}

func (c *Config) SetForFolders() {
	folderToScan, err := getFoldersToScan(c.Path.PathToScan)
	PrintError("Failed to Read Directory", err)
	c.Folder.FoldersToScan = folderToScan
}
