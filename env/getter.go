package env

import "os"

func getCurrentWorkingDirectory() (string, error) {
	return os.Getwd()
}

func getFoldersToScan(directoryToScan string) ([]os.DirEntry, error) {
	return os.ReadDir(directoryToScan)
}
