package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Args struct {
	BasePath string
}

func LoadArgs() (Args, error) {
	var args Args

	currentDir, err := os.Getwd()
	if err != nil {
		return args, fmt.Errorf("error getting current dir: %w", err)
	}

	flag.StringVar(&args.BasePath, "p", currentDir, "directory to cacluate stats for")
	flag.Parse()

	args.BasePath, err = filepath.Abs(args.BasePath)
	if err != nil {
		return args, fmt.Errorf("cannot find the directory provided!, %w", err)
	}

	return args, nil
}
