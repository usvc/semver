package repo

import (
	"fmt"
	"sort"

	"github.com/usvc/go-semver"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func TagCurrentGitCommit(withTag string, pathToRepository string) error {
	repository, err := git.PlainOpen(pathToRepository)
	if err != nil {
		return fmt.Errorf("failed to open the directory '%s' as a git repository: '%s'", pathToRepository, err)
	}
	head, err := repository.Head()
	if err != nil {
		return fmt.Errorf("failed to get HEAD of git repository: '%s'", err)
	}
	headCommitHash := head.Hash()
	_, err = repository.CreateTag(withTag, headCommitHash, nil)
	return err
}

func GetLatestSemverFromGitRepository(pathToRepository string) (*semver.Semver, error) {
	repository, err := git.PlainOpen(pathToRepository)
	if err != nil {
		return nil, fmt.Errorf("failed to open the directory '%s' as a git repository: '%s'", pathToRepository, err)
	}
	tags, err := repository.Tags()
	if err != nil {
		return nil, fmt.Errorf("failed to get tags from git repository at '%s': '%s'", pathToRepository, err)
	}
	var versions semver.Semvers
	tags.ForEach(func(tag *plumbing.Reference) error {
		tagName := tag.Name().Short()
		if semver.IsValid(tagName) {
			versions = append(versions, semver.Parse(tagName))
		}
		return nil
	})
	if versions.Len() == 0 {
		return &semver.Semver{
			Prefix: "v",
		}, nil
	}
	sort.Sort(versions)
	return versions[versions.Len()-1], nil
}
