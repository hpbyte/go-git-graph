package main

import (
	"fmt"
)

func main() {
	args := LoadArgs()
	fmt.Println("Welcome to GoGitGraph!, your chosen path: ", args.basePath)

	//res := ScanGitFolders(basePath)

	//Cache(res)

	gitConfig := LoadGitConfig()
	fmt.Println("git config email: ", gitConfig.User.Email)
	fmt.Println("git config name: ", gitConfig.User.Name)

	CalculateStats()
}
