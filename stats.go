package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
)

type Stats interface {
	Calculate() map[int]int
}

type ContributionStats struct {
	gitConfig *GitConfig
	stats     map[string]int
}

func (cs *ContributionStats) aggregate(path string) {
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
		if c.Author.Email == cs.gitConfig.User.Email {
			year, month, day := c.Committer.When.Date()

			key := fmt.Sprintf("%d-%02d-%02d", year, month, day)
			cs.stats[key] += 1
		}

		return nil
	})
}

func (cs ContributionStats) Calculate() map[string]int {
	repos := []string{"/Users/htoopyaelwin/Personal/adventofcode2023/.git"}

	for _, repo := range repos {
		cs.aggregate(repo)
	}

	return cs.stats
}
