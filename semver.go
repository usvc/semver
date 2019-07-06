package semver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const InvalidNumber = -1
const InvalidString = ""

// New is a convenience method that returns a Semver
// structure given a semver string, panicks if the provided
// string is not semver-formatted
func New(semver string) Semver {
	output := Semver{}
	if !IsValid(semver) {
		panic(fmt.Errorf("'%s' does not follow semver", semver))
	}
	output.FromString(semver)
	return output
}

// Semver stores a semver version in a structured format
// and provides methods for its manipulation
type Semver struct {
	Prefix         string
	Major          int
	Minor          int
	Patch          int
	Label          string
	LabelIteration int
}

// BumpMajor bumps the major version.
func (semver *Semver) BumpMajor() {
	semver.Major += 1
	semver.Minor = 0
	semver.Patch = 0
	semver.Label = InvalidString
	semver.LabelIteration = InvalidNumber
}

// BumpMinor bumps the minor version.
func (semver *Semver) BumpMinor() {
	semver.Minor += 1
	semver.Patch = 0
	semver.Label = InvalidString
	semver.LabelIteration = InvalidNumber
}

// BumpPatch bumps the patch version.
func (semver *Semver) BumpPatch() {
	semver.Patch += 1
	semver.Label = InvalidString
	semver.LabelIteration = InvalidNumber
}

// BumpLabel bumps the iteration number of the label
// iff there exists a label. If not label iteration
// number exists, it will be set to 0.
func (semver *Semver) BumpLabel() {
	if semver.Label != InvalidString {
		semver.LabelIteration += 1
	}
}

// FromString takes in an input string and assigns it
// to the appropriate fields in this instance of Semver
func (semver *Semver) FromString(input string) {
	if !IsValid(input) {
		panic(fmt.Errorf("'%s' does not follow semver", input))
	}
	matches := regexp.
		MustCompile(semverRegex).
		FindStringSubmatch(input)
	semver.Prefix = matches[1]
	var err error
	semver.Major, err = strconv.Atoi(matches[2])
	if err != nil {
		semver.Major = InvalidNumber
	}
	semver.Minor, err = strconv.Atoi(matches[3])
	if err != nil {
		semver.Minor = InvalidNumber
	}
	semver.Patch, err = strconv.Atoi(matches[4])
	if err != nil {
		semver.Patch = InvalidNumber
	}
	semver.Label = strings.Trim(matches[5], "-")
	semver.LabelIteration, err = strconv.Atoi(strings.Trim(matches[6], "."))
	if err != nil {
		semver.LabelIteration = InvalidNumber
	}
}

// HasLabel indicates if the label value exists logically
func (semver Semver) HasLabel() bool {
	return len(semver.Label) > 0
}

// HasLabelIteration indicates if there is a label iteration number
func (semver Semver) HasLabelIteration() bool {
	return semver.LabelIteration > InvalidNumber
}

func (semver Semver) IsGreaterThan(challenger Semver) bool {
	switch {
	case semver.Major > challenger.Major:
		return true
	case semver.Major < challenger.Major:
		return false
	case semver.Minor > challenger.Minor:
		return true
	case semver.Minor < challenger.Minor:
		return false
	case semver.Patch > challenger.Patch:
		return true
	case semver.Patch < challenger.Patch:
		return false
	case semver.Label == challenger.Label:
		if semver.LabelIteration > challenger.LabelIteration {
			return true
		} else {
			return false
		}
	}
	return false
}

// ToString returns a string representation of this Semver instance
func (semver Semver) ToString() string {
	output := fmt.Sprintf(
		"%s%v.%v.%v",
		semver.Prefix,
		semver.Major,
		semver.Minor,
		semver.Patch,
	)
	if semver.HasLabel() {
		output = fmt.Sprintf("%s-%s", output, semver.Label)
		if semver.HasLabelIteration() {
			output = fmt.Sprintf("%s.%v", output, semver.LabelIteration)
		}
	}
	return output
}
