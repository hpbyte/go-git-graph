package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
)

const outOfRange = 99999
const daysInLastSixMonths = 183
const weeksInLastSixMonths = 26

type Stats interface {
	Calculate() map[int][]int
}

type ContributionStats struct {
	gitConfig *GitConfig
	stats     map[int][]int
}

func (cs ContributionStats) aggregate(path string) {
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
			fmt.Println(c.Committer.When)
		}

		return nil
	})
}

func (cs ContributionStats) Calculate() map[int][]int {
	repos := []string{"/Users/htoopyaelwin/Personal/adventofcode2023/.git"}

	for _, repo := range repos {
		cs.aggregate(repo)
	}

	return cs.stats
}
