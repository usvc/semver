package semver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const InvalidNumber = -1
const InvalidString = ""

// NewSemver is a convenience method that returns a Semver
// structure given a semver string, panicks if the provided
// string is not semver-formatted
func NewSemver(semver string) Semver {
	output := Semver{}
	if !IsSemver(semver) {
		panic(fmt.Errorf("'%s' does not follow semver", semver))
	}
	output.FromString(semver)
	return output
}

type Semver struct {
	Prefix         string
	Major          int
	Minor          int
	Patch          int
	Label          string
	LabelIteration int
}

func (semver *Semver) BumpMajor() {
	semver.Major += 1
	semver.Minor = 0
	semver.Patch = 0
	semver.Label = InvalidString
	semver.LabelIteration = InvalidNumber
}

func (semver *Semver) BumpMinor() {
	semver.Minor += 1
	semver.Patch = 0
	semver.Label = InvalidString
	semver.LabelIteration = InvalidNumber
}

func (semver *Semver) BumpPatch() {
	semver.Patch += 1
	semver.Label = InvalidString
	semver.LabelIteration = InvalidNumber
}

func (semver *Semver) BumpLabel() {
	if semver.Label != InvalidString {
		semver.LabelIteration += 1
	}
}

func (semver *Semver) FromString(input string) {
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

func (semver Semver) HasLabel() bool {
	return len(semver.Label) > 0
}

func (semver Semver) HasLabelIteration() bool {
	return semver.LabelIteration > InvalidNumber
}

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
