package stats

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	. "go-git-graph/pkg/config"
)

type Stats interface {
	Calculate() map[int]int
}

type ContributionStats struct {
	Year      int
	GitConfig *GitConfig
	stats     map[string]int
}

func (cs *ContributionStats) aggregate(path string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("[Err]: opening the repo: %s\n", path)
	}

	headRef, err := repo.Head()
	if err != nil {
		return fmt.Errorf("[Err]: getting the HEAD of the repo: %s\n", path)
	}

	commits, err := repo.Log(&git.LogOptions{From: headRef.Hash()})
	if err != nil {
		return fmt.Errorf("[Err]: getting commit history: %s\n", path)
	}

	err = commits.ForEach(func(c *object.Commit) error {
		if c.Author.Email == cs.GitConfig.User.Email {
			year, month, day := c.Committer.When.Date()

			if year == cs.Year {
				key := fmt.Sprintf("%d-%02d-%02d", year, month, day)
				cs.stats[key] += 1
			}
		}

		return nil
	})

	return err
}

func (cs *ContributionStats) Calculate(repos map[string][]string) map[string]int {
	cs.stats = map[string]int{}

	log.Printf("[Log]: calculating for the year: %d...\n", cs.Year)

	for _, repoList := range repos {
		for _, repo := range repoList {
			err := cs.aggregate(repo)
			if err != nil {
				log.Println(err)
			}
		}
	}

	return cs.stats
}
