package git

import (
	"os"
	"path"

	git "gopkg.in/src-d/go-git.v4"
	plumbing "gopkg.in/src-d/go-git.v4/plumbing"
)

func GetTags(pathToRepository string) ([]string, error) {
	absolutePath := resolveAbsolutePath(pathToRepository)
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
	absolutePath := resolveAbsolutePath(pathToCheck)
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

func getCurrentWorkingDirectory() string {
	if workingDirectory, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		return workingDirectory
	}
}

func resolveAbsolutePath(providedPath string) string {
	if path.IsAbs(providedPath) {
		return providedPath
	}
	workingDirectory := getCurrentWorkingDirectory()
	resolvedPath := path.Join(workingDirectory, providedPath)
	return resolvedPath
}
