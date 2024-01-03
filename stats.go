package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
)

func aggregateStats(path string, stats map[int][]int) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		log.Fatalf("failed to open the repo: %s\n", err)
	}

	headRef, err := repo.Head()
	if err != nil {
		log.Fatalf("failed to get the HEAD of the repo: %s\n", err)
	}

	commits, err := repo.Log(&git.LogOptions{From: headRef.Hash()})
	if err != nil {
		log.Fatalf("failed to get commit history: %s\n", err)
	}

	err = commits.ForEach(func(c *object.Commit) error {
		fmt.Println(c.Committer.When)
		return nil
	})
}

func CalculateStats() {
	repos := []string{"/Users/htoopyaelwin/Personal/adventofcode2023/.git"}
	stats := make(map[int][]int)

	for _, repo := range repos {
		aggregateStats(repo, stats)
	}
}
