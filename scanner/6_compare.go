package scanner

import (
	"path/filepath"
	"strings"
)

func CompareFiles(files1, files2 []string) ([]string, []string, error) {
	root1 := commonDir(files1)
	root2 := commonDir(files2)

	set1 := make(map[string]struct{}, len(files1))
	for _, p := range files1 {
		rel := toRelSlash(p, root1)
		set1[strings.ToLower(rel)] = struct{}{}
	}

	set2 := make(map[string]struct{}, len(files2))
	for _, p := range files2 {
		rel := toRelSlash(p, root2)
		set2[strings.ToLower(rel)] = struct{}{}
	}

	onlyIn1 := diffKeys(set1, set2)
	onlyIn2 := diffKeys(set2, set1)
	return onlyIn1, onlyIn2, nil
}

func commonDir(paths []string) string {
	if len(paths) == 0 {
		return ""
	}

	split := func(s string) []string {
		clean := filepath.Clean(s)
		return strings.Split(filepath.Dir(clean), string(filepath.Separator))
	}

	parts := split(paths[0])

	for _, p := range paths[1:] {
		q := split(p)
		n := min(len(parts), len(q))

		i := 0
		for i < n && strings.EqualFold(parts[i], q[i]) {
			i++
		}
		parts = parts[:i]
		if len(parts) == 0 {
			break
		}
	}

	if len(parts) == 0 {
		vol := filepath.VolumeName(paths[0])
		if vol != "" {
			return vol + string(filepath.Separator)
		}
		return ""
	}
	return strings.Join(parts, string(filepath.Separator))
}

func toRelSlash(absPath, root string) string {
	if root == "" {
		return filepath.ToSlash(absPath)
	}
	rel, err := filepath.Rel(root, absPath)
	if err != nil {
		rel = absPath
	}
	return filepath.ToSlash(rel)
}

func diffKeys(a, b map[string]struct{}) []string {
	out := make([]string, 0, len(a))
	for k := range a {
		if _, ok := b[k]; !ok {
			out = append(out, k)
		}
	}

	sortStrings(out)
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
