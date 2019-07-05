package semver

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, &UtilsTests{})
}

func (s *UtilsTests) TestIsSemver() {
	// base cases
	s.True(IsSemver("1.2.3"))
	s.True(IsSemver("1.2.3-label"))
	s.True(IsSemver("1.2.3-label.4"))
	// prefixed cases
	s.True(IsSemver("v1.2.3"))
	s.True(IsSemver("v1.2.3-label"))
	s.True(IsSemver("v1.2.3-label.4"))
	s.True(IsSemver("V1.2.3"))
	s.True(IsSemver("V1.2.3-label"))
	s.True(IsSemver("V1.2.3-label.4"))
	s.True(IsSemver("V1.2.3-label_2.4"))
	s.True(IsSemver("V1.2.3-label_label.4"))
	// multidigit cases
	s.True(IsSemver("12.23.34"))
	s.True(IsSemver("12.23.34-label"))
	s.True(IsSemver("12.23.34-label.4"))
	// obvious errors
	s.False(IsSemver("a"))
	s.False(IsSemver("label"))
	s.False(IsSemver("1"))
	s.False(IsSemver("1-label"))
	s.False(IsSemver("1-label.2"))
	s.False(IsSemver("1.2"))
	s.False(IsSemver("1.2-label"))
	s.False(IsSemver("1.2-label.3"))
	// edge errors
	s.False(IsSemver("vv1.2.3"))
	s.False(IsSemver("Vv1.2.3"))
	s.False(IsSemver("vV1.2.3"))
	s.False(IsSemver("v1.2.3label"))
	s.False(IsSemver("v1.2.3_label"))
	s.False(IsSemver("v1.2.3-label.4.4"))
	s.False(IsSemver("v1.2.3-label.4-4"))
}
