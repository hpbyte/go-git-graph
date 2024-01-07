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

	fmt.Println("Welcome to GoGitGraph!, your chosen path: ", config.BasePath)

	//res := GitFolderScanner{}.Scan(args.BasePath)

	//Cacher{}.Create(res)

	contributionStats := ContributionStats{gitConfig: &config.GitConfig, stats: map[string]int{}}.Calculate()
	fmt.Println("calculated contributionStats: ", contributionStats)

	contributionChart := ContributionChart{Data: contributionStats, Year: 2024}
	contributionChart.Render()
}
