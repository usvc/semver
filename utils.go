package semver

import (
	"regexp"
)

const semverRegex = `^(?P<Prefix>[vV]?)` +
	`(?P<Major>[0-9]+)\.` +
	`(?P<Minor>[0-9]+)\.` +
	`(?P<Patch>[0-9]+)` +
	`(?P<Label>[\-][a-zA-Z0-9\_\-]+)?(?P<LabelIteration>\.[0-9]+)?$`

func IsValid(semver string) bool {
	matched, err := regexp.Match(semverRegex, []byte(semver))
	if err != nil {
		panic(err)
	}
	return matched
}
