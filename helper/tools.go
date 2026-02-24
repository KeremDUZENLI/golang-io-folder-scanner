package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

var chunkRegex = regexp.MustCompile(`(\d+|\D+)`)

func CanonicalPath(base string) string {
	pathAbs, _ := filepath.Abs(base)
	clean := filepath.Clean(pathAbs)
	return filepath.ToSlash(clean)
}

func RelativePath(base, path string) string {
	rel, err := filepath.Rel(base, path)
	if err != nil {
		return path
	}
	return filepath.ToSlash(rel)
}

func PrintError(err error) {
	if err != nil {
		fmt.Printf("\n[ERROR] %v\n", err)
		os.Exit(1)
	}
}

func SortEntries(entries []os.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return naturalLess(entries[i].Name(), entries[j].Name())
	})
}

func naturalLess(a, b string) bool {
	aChunks := chunkRegex.FindAllString(a, -1)
	bChunks := chunkRegex.FindAllString(b, -1)

	for i := 0; i < len(aChunks) && i < len(bChunks); i++ {
		aChunk, bChunk := aChunks[i], bChunks[i]
		aNum, aErr := strconv.Atoi(aChunk)
		bNum, bErr := strconv.Atoi(bChunk)

		if aErr == nil && bErr == nil {
			if aNum != bNum {
				return aNum < bNum
			}
		} else if aChunk != bChunk {
			return aChunk < bChunk
		}
	}
	return len(aChunks) < len(bChunks)
}
