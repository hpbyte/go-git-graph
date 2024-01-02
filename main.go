package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	var path string
	flag.StringVar(&path, "p", currentDir, "directory to cacluate stats for")
	flag.Parse()

	basePath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Welcome to GoGitGraph!, your chosen path: ", basePath)
}
