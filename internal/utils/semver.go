package utils

import (
	"fmt"
	"strings"

	"github.com/usvc/go-semver"
)

func GetSemverFromArguments(args []string) (*semver.Semver, error) {
	versionFromArguments := strings.Join(args, ".")
	if !semver.IsValid(versionFromArguments) {
		return nil, fmt.Errorf("parsed input '%s' is not a valid semver string", versionFromArguments)
	}
	return semver.Parse(versionFromArguments), nil
}
