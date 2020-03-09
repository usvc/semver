package utils

import (
	"fmt"
	"os"

	"github.com/usvc/go-semver"
	"gitlab.com/usvc/utils/semver/internal/repo"
)

func GetCurrentRepositoryLatestSemver() (*semver.Semver, error) {
	var (
		currentDirectory string
		err              error
		version          *semver.Semver
	)
	currentDirectory, err = os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: '%s'", err)
	}
	version, err = repo.GetLatestSemverFromGitRepository(currentDirectory)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve latest semver tag: '%s'", err)
	}
	return version, nil
}

func SetCurrentRepositorySemver(version *semver.Semver) error {
	currentDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to getting current working directory: '%s'", err)
	}
	err = repo.TagCurrentGitCommit(version.String(), currentDirectory)
	if err != nil {
		return fmt.Errorf("failed to add tag '%s' to repository at '%s': '%s'", version.String(), currentDirectory, err)
	}
	return nil
}
