package cache

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Cacher struct{}

func (c Cacher) getCacheFilePath() string {
	const cacheFileName = ".go-git-graph.cache.json"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicf("[Err]: loading home dir: %s\n", err)
	}

	cacheFilePath := filepath.Join(homeDir, cacheFileName)

	return cacheFilePath
}

func (c Cacher) createCacheFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("[Err]: creating cache file: %s\n", err)
	}
	defer file.Close()

	log.Println("[Log]: new cache file created.")
}

func (c Cacher) getOrCreateCacheFile() string {
	cacheFilePath := c.getCacheFilePath()

	if _, err := os.Stat(cacheFilePath); os.IsNotExist(err) {
		c.createCacheFile(cacheFilePath)
	}

	return cacheFilePath
}

func (c Cacher) Create(data map[string][]string) map[string][]string {
	path := c.getOrCreateCacheFile()
	// convert to json
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("[Err]: marshaling data to JSON: %v\n", err)
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		log.Fatalf("[Err]: writing cache results: %v\n", err)
	}

	log.Println("[Log]: successfully cached results")

	return data
}

func (c Cacher) Fetch() map[string][]string {
	path := c.getOrCreateCacheFile()

	byteData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string][]string
	json.Unmarshal([]byte(byteData), &data)

	return data
}

func (c Cacher) Clear() {
	cacheFilePath := c.getCacheFilePath()
	err := os.Remove(cacheFilePath)
	if err != nil {
		log.Fatalf("[Err]: deleting cache: %s\n", err)
	}
}
