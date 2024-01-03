package main

import (
	"fmt"
	"log"
)

func main() {
	args, err := LoadArgs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Welcome to GoGitGraph!, your chosen path: ", args.BasePath)

	//res := GitFolderScanner{}.Scan(args.BasePath)

	//Cacher{}.Create(res)

	var gitConfig GitConfig
	if err := gitConfig.Load(); err != nil {
		log.Fatal(err)
	}

	contributionStats := ContributionStats{gitConfig: &gitConfig, stats: map[int][]int{}}.Calculate()

	fmt.Println("calculated contributionStats: ", contributionStats)
}
