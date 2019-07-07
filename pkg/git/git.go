package git

import (
	git "gopkg.in/src-d/go-git.v4"
	plumbing "gopkg.in/src-d/go-git.v4/plumbing"

	"gitlab.com/usvc/utils/semver/internal/utils"
)

func AddTag(pathToRepository string, tag string) error {
	absolutePath := utils.ResolveAbsolutePath(pathToRepository)
	repository, err := git.PlainOpenWithOptions(
		absolutePath,
		&git.PlainOpenOptions{
			DetectDotGit: true,
		},
	)
	if err != nil {
		return err
	}
	headRef, headErr := repository.Head()
	if headErr != nil {
		return err
	}
	hash := headRef.Hash()
	refName := plumbing.ReferenceName("refs/tags/" + tag)
	ref := plumbing.NewHashReference(refName, hash)
	err = repository.Storer.SetReference(ref)
	if err != nil {
		return err
	}

	return nil
}

func GetTags(pathToRepository string) ([]string, error) {
	absolutePath := utils.ResolveAbsolutePath(pathToRepository)
	repository, err := git.PlainOpenWithOptions(
		absolutePath,
		&git.PlainOpenOptions{
			DetectDotGit: true,
		},
	)
	if err != nil {
		return []string{}, err
	}
	ref, err := repository.Tags()
	if err != nil {
		return []string{}, err
	}
	var tags []string
	ref.ForEach(func(r *plumbing.Reference) error {
		tags = append(tags, r.Name().Short())
		return nil
	})
	return tags, nil
}

func IsGitRepository(pathToCheck string) (bool, error) {
	absolutePath := utils.ResolveAbsolutePath(pathToCheck)
	if _, err := git.PlainOpenWithOptions(
		absolutePath,
		&git.PlainOpenOptions{
			DetectDotGit: true,
		},
	); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
