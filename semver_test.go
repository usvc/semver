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

func (s *SemverTests) TestNewSemver() {
	s.Equal(NewSemver("1.2.3-label.1"), Semver{"", 1, 2, 3, "label", 1})
	s.Equal(NewSemver("10.20.30-label1.10"), Semver{"", 10, 20, 30, "label1", 10})
	s.Equal(NewSemver("v10.20.30-label1.10"), Semver{"v", 10, 20, 30, "label1", 10})
	s.Equal(NewSemver("V10.20.30-label1.10"), Semver{"V", 10, 20, 30, "label1", 10})
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
