package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

type Args struct {
	basePath string
}

func LoadArgs() Args {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var path string
	flag.StringVar(&path, "p", currentDir, "directory to cacluate stats for")
	flag.Parse()

	basePath, err := filepath.Abs(path)
	if err != nil {
		log.Panic("cannot find the directory provided!")
	}

	var args Args
	args.basePath = basePath

	return args
}
