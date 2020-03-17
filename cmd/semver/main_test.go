package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mainTest struct {
	suite.Suite
}

func Test_main(t *testing.T) {
	suite.Run(t, &mainTest{})
}

func (s *mainTest) TestGetCommand() {
	cmd := GetCommand()
	s.Equal("semver", cmd.Use)
	s.NotNil(cmd.Run)
	s.Len(cmd.Commands(), 1)
}
