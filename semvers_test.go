package semver

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SemversTest struct {
	suite.Suite
}

func TestSemvers(t *testing.T) {
	suite.Run(t, &SemversTest{})
}

func (s *SemversTest) TestSort() {
	var semvers Semvers = []Semver{
		{
			Major: 2,
			Minor: 1,
			Patch: 1,
		},
		{
			Major: 1,
			Minor: 3,
			Patch: 0,
		},
		{
			Major: 1,
			Minor: 2,
			Patch: 4,
		},
		{
			Major: 1,
			Minor: 2,
			Patch: 3,
		},
	}
	s.False(sort.IsSorted(semvers))
	sort.Sort(semvers)
	s.True(sort.IsSorted(semvers))
	s.Equal(semvers[0].Patch, 3)
	s.Equal(semvers[1].Patch, 4)
	s.Equal(semvers[2].Minor, 3)
	s.Equal(semvers[3].Major, 2)
}
