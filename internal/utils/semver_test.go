package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SemverTests struct {
	suite.Suite
}

func TestSemver(t *testing.T) {
	suite.Run(t, &SemverTests{})
}

func (s *SemverTests) TestGetSemverFromArguments() {
	semver, err := GetSemverFromArguments([]string{"1", "2", "3"})
	s.Nil(err)
	s.Equal("1.2.3", semver.String())
	semver, err = GetSemverFromArguments([]string{"10", "20", "30"})
	s.Nil(err)
	s.Equal("10.20.30", semver.String())
}

func (s *SemverTests) TestGetSemverFromArguments_error() {
	semver, err := GetSemverFromArguments([]string{"a", "2", "3"})
	s.Nil(semver)
	s.NotNil(err)
	s.Contains(err.Error(), "not a valid semver")
	semver, err = GetSemverFromArguments([]string{"1", "b", "3"})
	s.Nil(semver)
	s.NotNil(err)
	s.Contains(err.Error(), "not a valid semver")
	semver, err = GetSemverFromArguments([]string{"1", "2", "c"})
	s.Nil(semver)
	s.NotNil(err)
	s.Contains(err.Error(), "not a valid semver")
}
