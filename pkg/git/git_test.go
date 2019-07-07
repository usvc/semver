package git

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/usvc/utils/semver/internal/utils"
)

type GitTests struct {
	suite.Suite
}

func TestGit(t *testing.T) {
	suite.Run(t, &GitTests{})
}

func (s *GitTests) TestGetTags() {
	cwd := utils.GetCurrentWorkingDirectory()
	tags, err := GetTags(cwd)
	s.Nil(err)
	s.True(len(tags) > 0)
}

func (s *GitTests) TestIsRepository() {
	cwd := utils.GetCurrentWorkingDirectory()
	is, err := IsGitRepository(cwd)
	s.Nil(err)
	s.True(is)
}
