package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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

	res := ScanGitFolders(basePath)

	Cache(res)
}
