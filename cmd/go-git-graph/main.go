package main

import (
	"fmt"
	. "go-git-graph/pkg/cache"
	. "go-git-graph/pkg/chart"
	. "go-git-graph/pkg/config"
	. "go-git-graph/pkg/scan"
	. "go-git-graph/pkg/stats"
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
	if cache != nil && cache[config.BasePath] != nil {
		log.Println("[Log]: provided path has been scanned before, using cached repo lists...")
	} else {
		res := GitFolderScanner{}.Scan(config.BasePath)

		cache = cacher.Create(res)
	}

	contributionStats := ContributionStats{GitConfig: &config.GitConfig, Year: config.Year}
	stats := contributionStats.Calculate(cache)

	contributionChart := ContributionChart{Data: stats, Year: config.Year}
	contributionChart.Render()
}
