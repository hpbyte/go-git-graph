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
	cache := cacher.Fetch()
	if cache == nil {
		res := GitFolderScanner{}.Scan(config.BasePath)

		cache = cacher.Create(res)
	}

	contributionStats := ContributionStats{gitConfig: &config.GitConfig, stats: map[string]int{}}.Calculate(cache)

	contributionChart := ContributionChart{Data: contributionStats, Year: 2024}
	contributionChart.Render()
}
