package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// recursively scan .git folders in the given directory
func scan(basePath string, path string, scanned map[string][]string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			entryName := entry.Name()
			entryPath := filepath.Join(path, entryName)

			// found .git folder
			if strings.HasSuffix(entryName, ".git") {
				if _, err := os.Stat(entryPath); err == nil {
					scanned[basePath] = append(scanned[basePath], entryPath)
				}
			} else if entryName == "node_modules" || entryName == "vendor" {
				continue
			} else {
				scan(basePath, entryPath, scanned)
			}
		}
	}
}

// track the scanned .git folder paths
func ScanGitFolders(path string) map[string][]string {
	scanned := make(map[string][]string)

	scan(path, path, scanned)

	return scanned
}
