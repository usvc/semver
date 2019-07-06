package git

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GitTests struct {
	suite.Suite
}

func TestGit(t *testing.T) {
	suite.Run(t, &GitTests{})
}

func (s *GitTests) TestGetTags() {
	cwd := getCurrentWorkingDirectory()
	tags, err := GetTags(cwd)
	s.Nil(err)
	s.True(len(tags) > 0)
}

func (s *GitTests) TestIsRepository() {
	cwd := getCurrentWorkingDirectory()
	is, err := IsGitRepository(cwd)
	s.Nil(err)
	s.True(is)
}
