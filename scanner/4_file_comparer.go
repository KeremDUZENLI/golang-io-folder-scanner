package scanner

import "strings"

func CompareFiles(files1, files2 []string) ([]string, []string) {
	root1 := commonDirCanonical(files1)
	root2 := commonDirCanonical(files2)

	rel1 := make([]string, 0, len(files1))
	for _, p := range files1 {
		rel1 = append(rel1, relCanonical(p, root1))
	}
	rel2 := make([]string, 0, len(files2))
	for _, p := range files2 {
		rel2 = append(rel2, relCanonical(p, root2))
	}

	set1 := make(map[string]struct{}, len(rel1))
	for _, r := range rel1 {
		set1[r] = struct{}{}
	}
	set2 := make(map[string]struct{}, len(rel2))
	for _, r := range rel2 {
		set2[r] = struct{}{}
	}

	onlyIn1 := make([]string, 0, len(rel1))
	seen1 := make(map[string]struct{}, 8)
	for _, r := range rel1 {
		if _, ok := set2[r]; !ok {
			if _, dup := seen1[r]; !dup {
				onlyIn1 = append(onlyIn1, r)
				seen1[r] = struct{}{}
			}
		}
	}

	onlyIn2 := make([]string, 0, len(rel2))
	seen2 := make(map[string]struct{}, 8)
	for _, r := range rel2 {
		if _, ok := set1[r]; !ok {
			if _, dup := seen2[r]; !dup {
				onlyIn2 = append(onlyIn2, r)
				seen2[r] = struct{}{}
			}
		}
	}

	return onlyIn1, onlyIn2
}

func commonDirCanonical(paths []string) string {
	if len(paths) == 0 {
		return ""
	}
	parts := strings.Split(paths[0], "/")
	for _, p := range paths[1:] {
		q := strings.Split(p, "/")
		n := min(len(parts), len(q))
		i := 0
		for i < n && parts[i] == q[i] {
			i++
		}
		parts = parts[:i]
		if len(parts) == 0 {
			break
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, "/")
}

func relCanonical(abs, root string) string {
	if root == "" {
		return abs
	}
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}
	if strings.HasPrefix(abs, root) {
		return strings.TrimPrefix(abs, root)
	}
	return abs
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
