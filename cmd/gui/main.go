package main

import (
	"fmt"
	"strings"

	"github.com/KeremDUZENLI/golang-io-folder-scanner/env"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/helper"
	"github.com/KeremDUZENLI/golang-io-folder-scanner/scanner"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// ---------- helpers ----------

func parseListCSV(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func buildGuides(ancestorLast []bool) string {
	var b strings.Builder
	for i := 0; i < len(ancestorLast)-1; i++ {
		if ancestorLast[i] {
			b.WriteString("    ")
		} else {
			b.WriteString("│   ")
		}
	}
	return b.String()
}

// ---------- backend operations ----------

func opListFolders(base string, extraSkip []string) string {
	cfg := env.ConfigDefault
	cfg.PathToScan = base

	foldersToSkipAll := append(cfg.FoldersToSkip, extraSkip...)

	folders := scanner.ListFolders(cfg.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)

	var b strings.Builder
	fmt.Fprintf(&b, "FILTERED FOLDERS (base = %s)\n\n", cfg.PathToScan)
	for _, f := range foldersFiltered {
		fmt.Fprintln(&b, helper.RelativePath(cfg.PathToScan, f))
	}
	fmt.Fprintf(&b, "\nTOTAL: %d\n", len(foldersFiltered))
	return b.String()
}

func opListFiles(base string, extraSkip, suffixes []string) string {
	cfg := env.ConfigDefault
	cfg.PathToScan = base
	if len(suffixes) > 0 {
		cfg.SuffixesToScan = suffixes
	}
	foldersToSkipAll := append(cfg.FoldersToSkip, extraSkip...)

	folders := scanner.ListFolders(cfg.PathToScan)
	foldersFiltered := scanner.FilterFolders(folders, foldersToSkipAll)
	files := scanner.ListFiles(foldersFiltered)
	filesFiltered := scanner.FilterFiles(files, cfg.SuffixesToScan)

	var b strings.Builder
	fmt.Fprintf(&b, "FILTERED FILES (base = %s)\n\n", cfg.PathToScan)
	for _, f := range filesFiltered {
		fmt.Fprintln(&b, helper.RelativePath(cfg.PathToScan, f))
	}
	fmt.Fprintf(&b, "\nTOTAL: %d\n", len(filesFiltered))
	return b.String()
}

func opScanFilesContent(base string, extraSkip, suffixes []string) string {
	cfg := env.ConfigDefault
	if len(suffixes) > 0 {
		cfg.SuffixesToScan = suffixes
	}

	skip := append(cfg.FoldersToSkip, extraSkip...)

	folders := scanner.FilterFolders(scanner.ListFolders(base), skip)
	files := scanner.FilterFiles(scanner.ListFiles(folders), cfg.SuffixesToScan)
	contents := scanner.ScanFilesContent(files)

	var b strings.Builder
	b.WriteString("CONTENT OF FILES (base = " + base + ")\n\n")

	sep := strings.Repeat("-", 100) + "\n"

	for _, c := range contents {
		rel := helper.RelativePath(base, c.Path)
		b.WriteString(rel + "=\n")
		b.WriteString(c.Content + "\n")
		b.WriteString(sep)
	}

	return b.String()
}

func opTree(base string, extraSkip, extraTreeSkip []string) string {
	cfg := env.ConfigDefault

	skip := append(cfg.FoldersToSkip, extraSkip...)
	treeSkip := append(cfg.FoldersTreeToSkip, extraTreeSkip...)

	folders := scanner.FilterFolders(scanner.ListFolders(base), skip)
	files := scanner.ListFiles(folders)
	items := scanner.CreateTree(folders, files, treeSkip)

	if len(items) == 0 {
		return "No items."
	}

	var b strings.Builder
	b.WriteString("ASCII TREE (base = " + base + ")\n\n")

	for _, item := range items[1:] {
		prefix := buildGuides(item.AncestorLast)
		branch := "├── "
		if len(item.AncestorLast) > 0 && item.AncestorLast[len(item.AncestorLast)-1] {
			branch = "└── "
		}
		rel := helper.RelativePath(base, item.Path)
		parts := strings.Split(rel, "/")
		name := parts[len(parts)-1]
		fmt.Fprintf(&b, "%s%s%s\n", prefix, branch, name)
	}

	return b.String()
}

func opEmptyFolders(base string, extraSkip []string) string {
	cfg := env.ConfigDefault
	skip := append(cfg.FoldersToSkip, extraSkip...)

	folders := scanner.FilterFolders(scanner.ListFolders(base), skip)
	empty := scanner.FindFoldersEmpty(folders)

	var b strings.Builder
	b.WriteString("EMPTY FOLDERS (base = " + base + ")\n\n")
	for _, f := range empty {
		b.WriteString(helper.RelativePath(base, f) + "\n")
	}
	b.WriteString(fmt.Sprintf("\nTOTAL: %d\n", len(empty)))
	return b.String()
}

func opFoldersBySuffix(base string, extraSkip, suffixes []string) string {
	cfg := env.ConfigDefault
	if len(suffixes) > 0 {
		cfg.SuffixesToScan = suffixes
	}
	skip := append(cfg.FoldersToSkip, extraSkip...)

	folders := scanner.FilterFolders(scanner.ListFolders(base), skip)
	found := scanner.FindFoldersByFileSuffix(folders, cfg.SuffixesToScan)

	var b strings.Builder
	b.WriteString("FOUND FOLDERS (base = " + base + ")\n\n")
	for _, f := range found {
		b.WriteString(helper.RelativePath(base, f) + "\n")
	}
	b.WriteString(fmt.Sprintf("\nTOTAL: %d\n", len(found)))
	return b.String()
}

func opCompare(base1, base2 string, extraSkip []string) string {
	cfg := env.ConfigDefault
	skip := append(cfg.FoldersToSkip, extraSkip...)

	folders1 := scanner.FilterFolders(scanner.ListFolders(base1), skip)
	folders2 := scanner.FilterFolders(scanner.ListFolders(base2), skip)

	files1 := scanner.ListFiles(folders1)
	files2 := scanner.ListFiles(folders2)

	only1, only2 := scanner.CompareFiles(base1, base2, files1, files2)

	var b strings.Builder
	b.WriteString("FILE COMPARISON\n")

	b.WriteString("\nONLY IN " + base1 + "\n\n")
	for _, p := range only1 {
		b.WriteString(helper.RelativePath(base1, p) + "\n")
	}
	b.WriteString(fmt.Sprintf("\nTOTAL: %d\n", len(only1)))

	b.WriteString("\n\nONLY IN " + base2 + "\n\n")
	for _, p := range only2 {
		b.WriteString(helper.RelativePath(base2, p) + "\n")
	}
	b.WriteString(fmt.Sprintf("\nTOTAL: %d\n", len(only2)))

	return b.String()
}

// ---------- main GUI ----------

func main() {
	a := app.New()
	w := a.NewWindow("IO Folder Scanner - GUI")

	currentPathEntry := widget.NewEntry()
	currentPathEntry.SetPlaceHolder("Base path (e.g. . or /Users/you/project)")

	browseBaseBtn := widget.NewButton("Browse…", func() {
		dlg := dialog.NewFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil || list == nil {
				return
			}
			currentPathEntry.SetText(list.Path())
		}, w)
		dlg.Show()
	})

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Output will appear here...")
	output.Wrapping = fyne.TextWrapWord

	copyBtn := widget.NewButton("Copy output", func() {
		w.Clipboard().SetContent(output.Text)
		dialog.ShowInformation("Copied", "Output copied to clipboard!", w)
	})

	// ===== Tab 1: List folders =====

	listFoldersSkipEntry := widget.NewEntry()
	listFoldersSkipEntry.SetPlaceHolder("Extra folders to skip (comma separated)")

	listFoldersRun := widget.NewButton("Run", func() {
		base := strings.TrimSpace(currentPathEntry.Text)
		if base == "" {
			base = "."
		}
		base = helper.CanonicalPath(base)
		extraSkip := parseListCSV(listFoldersSkipEntry.Text)
		res := opListFolders(base, extraSkip)
		output.SetText(res)
	})

	tab1Content := container.NewVBox(
		widget.NewLabel("1) List folders"),
		widget.NewLabel("Extra folders to skip:"),
		listFoldersSkipEntry,
		listFoldersRun,
	)

	// ===== Tab 2: List files =====

	listFilesSkipEntry := widget.NewEntry()
	listFilesSkipEntry.SetPlaceHolder("Extra folders to skip (comma separated)")

	listFilesSuffixesEntry := widget.NewEntry()
	listFilesSuffixesEntry.SetPlaceHolder("Suffixes to scan (comma separated, e.g. .go,.py,.js)")

	listFilesRun := widget.NewButton("Run", func() {
		base := strings.TrimSpace(currentPathEntry.Text)
		if base == "" {
			base = "."
		}
		base = helper.CanonicalPath(base)
		extraSkip := parseListCSV(listFilesSkipEntry.Text)
		suffixes := parseListCSV(listFilesSuffixesEntry.Text)
		res := opListFiles(base, extraSkip, suffixes)
		output.SetText(res)
	})

	tab2Content := container.NewVBox(
		widget.NewLabel("2) List files"),
		widget.NewLabel("Extra folders to skip:"),
		listFilesSkipEntry,
		widget.NewLabel("Suffixes to scan:"),
		listFilesSuffixesEntry,
		listFilesRun,
	)

	// ===== Tab 3: Scan content of files =====

	scanSkipEntry := widget.NewEntry()
	scanSkipEntry.SetPlaceHolder("Extra folders to skip (comma separated)")

	scanSuffixesEntry := widget.NewEntry()
	scanSuffixesEntry.SetPlaceHolder("Suffixes to scan (comma separated, e.g. .go,.py,.js)")

	scanRun := widget.NewButton("Run", func() {
		base := strings.TrimSpace(currentPathEntry.Text)
		if base == "" {
			base = "."
		}
		base = helper.CanonicalPath(base)
		extraSkip := parseListCSV(scanSkipEntry.Text)
		suffixes := parseListCSV(scanSuffixesEntry.Text)

		// This is the heavy one; still running inline to behave like CLI.
		res := opScanFilesContent(base, extraSkip, suffixes)
		output.SetText(res)
	})

	tab3Content := container.NewVBox(
		widget.NewLabel("3) Scan content of files"),
		widget.NewLabel("Extra folders to skip:"),
		scanSkipEntry,
		widget.NewLabel("Suffixes to scan:"),
		scanSuffixesEntry,
		scanRun,
	)

	// ===== Tab 4: Create ASCII tree =====

	treeSkipEntry := widget.NewEntry()
	treeSkipEntry.SetPlaceHolder("Extra folders to skip (comma separated)")

	treeTreeSkipEntry := widget.NewEntry()
	treeTreeSkipEntry.SetPlaceHolder("Extra folders-tree to skip (comma separated, e.g. img,images)")

	treeRun := widget.NewButton("Run", func() {
		base := strings.TrimSpace(currentPathEntry.Text)
		if base == "" {
			base = "."
		}
		base = helper.CanonicalPath(base)
		extraSkip := parseListCSV(treeSkipEntry.Text)
		extraTreeSkip := parseListCSV(treeTreeSkipEntry.Text)
		res := opTree(base, extraSkip, extraTreeSkip)
		output.SetText(res)
	})

	tab4Content := container.NewVBox(
		widget.NewLabel("4) Create ASCII tree"),
		widget.NewLabel("Extra folders to skip:"),
		treeSkipEntry,
		widget.NewLabel("Extra tree folders to skip:"),
		treeTreeSkipEntry,
		treeRun,
	)

	// ===== Tab 5: Find empty folders =====

	emptySkipEntry := widget.NewEntry()
	emptySkipEntry.SetPlaceHolder("Extra folders to skip (comma separated)")

	emptyRun := widget.NewButton("Run", func() {
		base := strings.TrimSpace(currentPathEntry.Text)
		if base == "" {
			base = "."
		}
		base = helper.CanonicalPath(base)
		extraSkip := parseListCSV(emptySkipEntry.Text)
		res := opEmptyFolders(base, extraSkip)
		output.SetText(res)
	})

	tab5Content := container.NewVBox(
		widget.NewLabel("5) Find empty folders"),
		widget.NewLabel("Extra folders to skip:"),
		emptySkipEntry,
		emptyRun,
	)

	// ===== Tab 6: Find folders by file suffix =====

	ffSkipEntry := widget.NewEntry()
	ffSkipEntry.SetPlaceHolder("Extra folders to skip (comma separated)")

	ffSuffixesEntry := widget.NewEntry()
	ffSuffixesEntry.SetPlaceHolder("Suffixes to scan (comma separated, e.g. .go,.py)")

	ffRun := widget.NewButton("Run", func() {
		base := strings.TrimSpace(currentPathEntry.Text)
		if base == "" {
			base = "."
		}
		base = helper.CanonicalPath(base)
		extraSkip := parseListCSV(ffSkipEntry.Text)
		suffixes := parseListCSV(ffSuffixesEntry.Text)
		res := opFoldersBySuffix(base, extraSkip, suffixes)
		output.SetText(res)
	})

	tab6Content := container.NewVBox(
		widget.NewLabel("6) Find folders by file suffix"),
		widget.NewLabel("Extra folders to skip:"),
		ffSkipEntry,
		widget.NewLabel("Suffixes to scan:"),
		ffSuffixesEntry,
		ffRun,
	)

	// ===== Tab 7: Compare two paths =====

	compareBase2Entry := widget.NewEntry()
	compareBase2Entry.SetPlaceHolder("Second path (base2)")

	compareBrowseBtn := widget.NewButton("Browse…", func() {
		dlg := dialog.NewFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil || list == nil {
				return
			}
			compareBase2Entry.SetText(list.Path())
		}, w)
		dlg.Show()
	})

	compareSkipEntry := widget.NewEntry()
	compareSkipEntry.SetPlaceHolder("Extra folders to skip (comma separated)")

	compareRun := widget.NewButton("Run", func() {
		base1 := strings.TrimSpace(currentPathEntry.Text)
		if base1 == "" {
			base1 = "."
		}
		base1 = helper.CanonicalPath(base1)

		base2 := strings.TrimSpace(compareBase2Entry.Text)
		if base2 == "" {
			output.SetText("Please choose a second path (base2).")
			return
		}
		base2 = helper.CanonicalPath(base2)

		extraSkip := parseListCSV(compareSkipEntry.Text)
		res := opCompare(base1, base2, extraSkip)
		output.SetText(res)
	})

	tab7Content := container.NewVBox(
		widget.NewLabel("7) Compare two paths"),
		widget.NewLabel("Second path (base2):"),
		container.NewHBox(compareBase2Entry, compareBrowseBtn),
		widget.NewLabel("Extra folders to skip:"),
		compareSkipEntry,
		compareRun,
	)

	// ===== Tabs container =====

	tabs := container.NewAppTabs(
		container.NewTabItem("1) Folders", tab1Content),
		container.NewTabItem("2) Files", tab2Content),
		container.NewTabItem("3) Scan", tab3Content),
		container.NewTabItem("4) Tree", tab4Content),
		container.NewTabItem("5) Empty", tab5Content),
		container.NewTabItem("6) By Suffix", tab6Content),
		container.NewTabItem("7) Compare", tab7Content),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// ===== Layout =====

	topBar := container.NewBorder(
		nil, nil,
		widget.NewLabel("Current path:"),
		browseBaseBtn,
		currentPathEntry,
	)

	leftSide := container.NewBorder(
		topBar,
		nil,
		nil,
		nil,
		tabs,
	)
	rightSide := container.NewBorder(
		container.NewHBox(copyBtn), // top
		nil,
		nil,
		nil,
		container.NewVScroll(output),
	)

	mainSplit := container.NewHSplit(
		leftSide,
		rightSide,
	)

	w.SetContent(mainSplit)
	w.Resize(fyne.NewSize(1200, 700))
	w.ShowAndRun()
}
