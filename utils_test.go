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

func (s *UtilsTests) TestIsValid() {
	// base cases
	s.True(IsValid("1.2.3"))
	s.True(IsValid("1.2.3-label"))
	s.True(IsValid("1.2.3-label.4"))
	// prefixed cases
	s.True(IsValid("v1.2.3"))
	s.True(IsValid("v1.2.3-label"))
	s.True(IsValid("v1.2.3-label.4"))
	s.True(IsValid("V1.2.3"))
	s.True(IsValid("V1.2.3-label"))
	s.True(IsValid("V1.2.3-label.4"))
	s.True(IsValid("V1.2.3-label_2.4"))
	s.True(IsValid("V1.2.3-label_label.4"))
	// multidigit cases
	s.True(IsValid("12.23.34"))
	s.True(IsValid("12.23.34-label"))
	s.True(IsValid("12.23.34-label.4"))
	// obvious errors
	s.False(IsValid("a"))
	s.False(IsValid("label"))
	s.False(IsValid("1"))
	s.False(IsValid("1-label"))
	s.False(IsValid("1-label.2"))
	s.False(IsValid("1.2"))
	s.False(IsValid("1.2-label"))
	s.False(IsValid("1.2-label.3"))
	// edge errors
	s.False(IsValid("vv1.2.3"))
	s.False(IsValid("Vv1.2.3"))
	s.False(IsValid("vV1.2.3"))
	s.False(IsValid("v1.2.3label"))
	s.False(IsValid("v1.2.3_label"))
	s.False(IsValid("v1.2.3-label.4.4"))
	s.False(IsValid("v1.2.3-label.4-4"))
}
