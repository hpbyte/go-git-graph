package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Cacher struct{}

func (c Cacher) getCacheFilePath() (string, error) {
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

func (c Cacher) Create(data map[string][]string) {
	cacheFilePath, err := c.getCacheFilePath()
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
