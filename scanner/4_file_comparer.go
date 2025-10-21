package scanner

func CompareFiles(files1, files2 []string) ([]string, []string) {
	set1 := make(map[string]struct{}, len(files1))
	for _, p := range files1 {
		set1[p] = struct{}{}
	}
	set2 := make(map[string]struct{}, len(files2))
	for _, p := range files2 {
		set2[p] = struct{}{}
	}

	onlyIn1 := make([]string, 0, len(files1))
	for _, p := range files1 {
		if _, ok := set2[p]; !ok {
			onlyIn1 = append(onlyIn1, p)
		}
	}

	onlyIn2 := make([]string, 0, len(files2))
	for _, p := range files2 {
		if _, ok := set1[p]; !ok {
			onlyIn2 = append(onlyIn2, p)
		}
	}

	return onlyIn1, onlyIn2
}
