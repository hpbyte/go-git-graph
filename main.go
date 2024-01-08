package main

import (
	"fmt"
	"log"
)

func main() {
	var configLoader Config
	config, err := configLoader.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Welcome to GoGitGraph!, your chosen path: %s\n\n\n", config.BasePath)

	cacher := Cacher{}

	if config.ClearCache {
		cacher.Clear()
	}

	cache := cacher.Fetch()
	if cache != nil {
		log.Println("[Log]: provided path has been scanned before, using cached repo lists...")
	} else {
		res := GitFolderScanner{}.Scan(config.BasePath)

		cache = cacher.Create(res)
	}

	contributionStats := ContributionStats{gitConfig: &config.GitConfig, year: config.Year}
	stats := contributionStats.Calculate(cache)

	contributionChart := ContributionChart{data: stats, year: config.Year}
	contributionChart.Render()
}
