package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCacheFilePath() (string, error) {
	cacheFilePath := "./.cache.json"

	if _, err := os.Stat(cacheFilePath); os.IsNotExist(err) {
		file, err := os.Create(cacheFilePath)
		if err != nil {
			return "", fmt.Errorf("error creating cache file: %s\n", err)
		}
		defer file.Close()

		log.Printf("new cache file created.")
	}

	return cacheFilePath, nil
}

func cache(data map[string][]string) {
	cacheFilePath, err := getCacheFilePath()
	if err != nil {
		log.Fatal(err)
		return
	}

	// convert to json
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling data to JSON: %v", err)
	}

	err = os.WriteFile(cacheFilePath, jsonData, 0644)
	if err != nil {
		log.Fatalf("error writing cache results: %v", err)
	}

	log.Printf("successfully cached results")
}

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
func scanGitFolders(path string) map[string][]string {
	scanned := make(map[string][]string)

	scan(path, path, scanned)

	return scanned
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	var path string
	flag.StringVar(&path, "p", currentDir, "directory to cacluate stats for")
	flag.Parse()

	basePath, err := filepath.Abs(path)
	if err != nil {
		log.Panic("cannot find the directory provided!")
	}

	fmt.Println("Welcome to GoGitGraph!, your chosen path: ", basePath)

	res := scanGitFolders(basePath)

	cache(res)
}
