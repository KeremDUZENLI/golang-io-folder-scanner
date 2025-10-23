package scanner

import "github.com/KeremDUZENLI/golang-io-folder-scanner/helper"

func CompareFiles(base1, base2 string, files1, files2 []string) ([]string, []string) {
	set1 := make(map[string]struct{}, len(files1))
	for _, p := range files1 {
		k := helper.RelativePath(base1, p)
		set1[k] = struct{}{}
	}
	set2 := make(map[string]struct{}, len(files2))
	for _, p := range files2 {
		k := helper.RelativePath(base2, p)
		set2[k] = struct{}{}
	}

	onlyIn1 := make([]string, 0, len(files1))
	for _, p := range files1 {
		k := helper.RelativePath(base1, p)
		if _, ok := set2[k]; !ok {
			onlyIn1 = append(onlyIn1, p)
		}
	}

	onlyIn2 := make([]string, 0, len(files2))
	for _, p := range files2 {
		k := helper.RelativePath(base2, p)
		if _, ok := set1[k]; !ok {
			onlyIn2 = append(onlyIn2, p)
		}
	}

	return onlyIn1, onlyIn2
}
