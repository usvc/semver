package semver

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

func (s *SemverTests) TestNew() {
	s.Equal(New("1.2.3-label.1"), Semver{"", 1, 2, 3, "label", 1})
	s.Equal(New("10.20.30-label1.10"), Semver{"", 10, 20, 30, "label1", 10})
	s.Equal(New("v10.20.30-label1.10"), Semver{"v", 10, 20, 30, "label1", 10})
	s.Equal(New("V10.20.30-label1.10"), Semver{"V", 10, 20, 30, "label1", 10})
}

func (s *SemverTests) TestSemverFromString_basic() {
	semver := Semver{}
	semver.FromString("1.2.3")
	s.Equal(semver.Prefix, InvalidString)
	s.Equal(semver.Major, 1)
	s.Equal(semver.Minor, 2)
	s.Equal(semver.Patch, 3)
	s.Equal(semver.Label, InvalidString)
	s.Equal(semver.LabelIteration, InvalidNumber)
}

func (s *SemverTests) TestSemverFromString_full() {
	semver := Semver{}
	semver.FromString("v1.2.3-label.4")
	s.Equal(semver.Prefix, "v")
	s.Equal(semver.Major, 1)
	s.Equal(semver.Minor, 2)
	s.Equal(semver.Patch, 3)
	s.Equal(semver.Label, "label")
	s.Equal(semver.LabelIteration, 4)
}

func (s *SemverTests) TestSemverIsGreaterThan() {
	// major
	semver1 := Semver{"v", 1, 2, 3, "", -1}
	semver2 := Semver{"v", 2, 0, 0, "", -1}
	s.False(semver1.IsGreaterThan(semver2))
	s.True(semver2.IsGreaterThan(semver1))
	semver2 = Semver{"v", 2, 3, 0, "", -1}
	s.False(semver1.IsGreaterThan(semver2))
	s.True(semver2.IsGreaterThan(semver1))
	semver2 = Semver{"v", 2, 0, 4, "", -1}
	s.False(semver1.IsGreaterThan(semver2))
	s.True(semver2.IsGreaterThan(semver1))
	// minor
	semver1 = Semver{"v", 1, 2, 3, "", -1}
	semver2 = Semver{"v", 1, 3, 0, "", -1}
	s.False(semver1.IsGreaterThan(semver2))
	s.True(semver2.IsGreaterThan(semver1))
	semver2 = Semver{"v", 1, 3, 4, "", -1}
	s.False(semver1.IsGreaterThan(semver2))
	s.True(semver2.IsGreaterThan(semver1))
	// patch
	semver1 = Semver{"v", 1, 2, 3, "", -1}
	semver2 = Semver{"v", 1, 2, 4, "", -1}
	s.False(semver1.IsGreaterThan(semver2))
	s.True(semver2.IsGreaterThan(semver1))
	// label
	semver1 = Semver{"v", 1, 2, 3, "label", 3}
	semver2 = Semver{"v", 1, 2, 3, "label", 5}
	s.False(semver1.IsGreaterThan(semver2))
	s.True(semver2.IsGreaterThan(semver1))
	semver2 = Semver{"v", 1, 2, 3, "label", 11}
	s.False(semver1.IsGreaterThan(semver2))
	s.True(semver2.IsGreaterThan(semver1))
}

func (s *SemverTests) TestSemverToString() {
	// basic test cases
	s.Equal(Semver{"", 1, 2, 3, "", -1}.ToString(), "1.2.3")
	s.Equal(Semver{"v", 1, 2, 3, "", -1}.ToString(), "v1.2.3")
	s.Equal(Semver{"V", 1, 2, 3, "", -1}.ToString(), "V1.2.3")
	s.Equal(Semver{"", 1, 2, 3, "label", -1}.ToString(), "1.2.3-label")
	s.Equal(Semver{"v", 1, 2, 3, "label", -1}.ToString(), "v1.2.3-label")
	s.Equal(Semver{"V", 1, 2, 3, "label", -1}.ToString(), "V1.2.3-label")
	s.Equal(Semver{"V", 1, 2, 3, "label", 4}.ToString(), "V1.2.3-label.4")
	s.Equal(Semver{"", 1, 2, 3, "label", 4}.ToString(), "1.2.3-label.4")

	// if label iteration exists without a label, ignore it
	s.Equal(Semver{"", 1, 2, 3, "", 4}.ToString(), "1.2.3")
}
