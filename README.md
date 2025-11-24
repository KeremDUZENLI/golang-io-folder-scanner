# golang-io-folder-scanner

A fast CLI to scan directories, read file contents, print ASCII trees, find files, and compare paths.

## Features

- 0 List folders & files
- 1 Filter folders
- 2 Filter files
- 3 Scan files content
- 4 Create tree ASCII
- 5 Find folders empty
- 6 Find folders by file suffix
- 7 Compare files in 2 different paths

## Downloads

You can download precompiled executables from the [Releases](https://github.com/KeremDUZENLI/golang-io-folder-scanner/releases/latest) page.

| Platform    | File                                                                                                                         |
| ----------- | ---------------------------------------------------------------------------------------------------------------------------- |
| **Linux**   | [IOScanner_Linux](https://github.com/KeremDUZENLI/golang-io-folder-scanner/releases/latest/download/IOScanner_Linux)         |
| **macOS**   | [IOScanner_MacOS](https://github.com/KeremDUZENLI/golang-io-folder-scanner/releases/latest/download/IOScanner_MacOS)         |
| **Windows** | [IOScanner_Windows](https://github.com/KeremDUZENLI/golang-io-folder-scanner/releases/latest/download/IOScanner_Windows.exe) |

## Menu

```
1) List folders
2) List files
3) Scan content of files
4) Create ASCII tree
5) Find folders empty
6) Find folders by file suffix
7) Compare two paths
(Press ENTER to quit)
```

## Inputs (prompts)

- **[NEW] Path to scan** — default: current working directory
- **[NEW] Suffixes to scan** — comma-separated (e.g., `.go, .md, .yml`)
- **[ADD] Folders to skip** — comma-separated (e.g., `__pycache__, node_modules, .git`)
- **[ADD] Folders tree to skip** — comma-separated (e.g., `img, images`)

> Tips
>
> - Paths are case-sensitive.
> - Folders and suffixes are case-insensitive.
> - [NEW] tag indicates overwritten value by input.
> - [ADD] tag indicates values to add by input.

## Output Snippets

**File Contents**

```
env/config.go=
package env
...

----------------------------------------------------------------------
env/setup.go=
package env
...
```

**ASCII Tree**

```
├── env
│   ├── config.go
│   └── setup.go
├── scanner
│   ├── 0_Lists.go
│   └── 7_CompareFiles.go
└── main.go
```

## LICENCE

This project is licensed under [**All Rights Reserved**](LICENSE).

© 2025 Kerem DÜZENLİ

## SUPPORT MY PROJECTS

If you find this resource valuable and would like to help support my education and doctoral research, please consider treating me to a cup of coffee (or tea) via Revolut.

<div align="center">
  <a href="https://revolut.me/krmdznl" target="_blank">
    <img src="https://img.shields.io/badge/Support%20My%20Projects-Donate%20via%20Revolut-orange?style=for-the-badge" alt="Support my projects via Revolut" />
  </a>
</div> <br>
