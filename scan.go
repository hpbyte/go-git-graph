package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Scanner interface {
	Scan() map[string][]string
}

type GitFolderScanner struct{}

// track the scanned .git folder paths
func (gfs GitFolderScanner) Scan(path string) map[string][]string {
	scanned := make(map[string][]string)

	// recursively scan .git folders in the given directory
	var scan func(basePath string, path string)

	scan = func(basePath string, path string) {
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
					scan(basePath, entryPath)
				}
			}
		}
	}

	scan(path, path)

	return scanned
}
