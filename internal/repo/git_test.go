package repo

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type GitTests struct {
	suite.Suite
}

func TestGit(t *testing.T) {
	suite.Run(t, &GitTests{})
}

func (s *GitTests) TestTagCurrentGitCommit() {
	expectedTag := "test_tag_current_git_commit"
	tempDir, err := ioutil.TempDir("", expectedTag)
	defer func() {
		os.RemoveAll(tempDir)
	}()
	s.Nil(err)
	repository, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL: "https://gitlab.com/usvc/utils/semver.git",
	})
	s.Nil(err)
	s.NotNil(repository)
	err = TagCurrentGitCommit(expectedTag, tempDir)
	s.Nil(err)
	tags, err := repository.Tags()
	s.Nil(err)
	var tagsAsStrings []string
	tags.ForEach(func(tag *plumbing.Reference) error {
		tagsAsStrings = append(tagsAsStrings, tag.Name().Short())
		return nil
	})
	s.Contains(tagsAsStrings, expectedTag)
}

func (s *GitTests) TestGetLatestSemverFromGitRepository() {
	directoryName := "test_get_latest_semver_from_git_repository"
	expectedTag := "123456.78.90"
	tempDir, err := ioutil.TempDir("", directoryName)
	defer func() {
		os.RemoveAll(tempDir)
	}()
	s.Nil(err)
	repository, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL: "https://gitlab.com/usvc/utils/semver.git",
	})
	s.Nil(err)
	s.NotNil(repository)
	err = TagCurrentGitCommit(expectedTag, tempDir)
	s.Nil(err)
	latest, err := GetLatestSemverFromGitRepository(tempDir)
	s.Nil(err)
	s.Equal(uint(123456), latest.VersionCore.Major)
	s.Equal(uint(78), latest.VersionCore.Minor)
	s.Equal(uint(90), latest.VersionCore.Patch)
}
