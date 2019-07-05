package semver

import (
	"os"
	"path"
	"regexp"

	git "gopkg.in/src-d/go-git.v4"
)

const semverRegex = `^(?P<Prefix>[vV]?)` +
	`(?P<Major>[0-9]+)\.` +
	`(?P<Minor>[0-9]+)\.` +
	`(?P<Patch>[0-9]+)` +
	`(?P<Label>[\-][a-zA-Z0-9\_\-]+)?(?P<LabelIteration>\.[0-9]+)?$`

func getCurrentWorkingDirectory() string {
	if workingDirectory, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		return workingDirectory
	}
}

func resolveAbsolutePath(providedPath string) string {
	if path.IsAbs(providedPath) {
		return providedPath
	}
	workingDirectory := getCurrentWorkingDirectory()
	resolvedPath := path.Join(workingDirectory, providedPath)
	return resolvedPath
}

func isGitRepository(pathToCheck string) bool {
	absolutePath := resolveAbsolutePath(pathToCheck)
	if _, err := git.PlainOpen(absolutePath); err != nil {
		return false
	} else {
		return true
	}
}

func IsSemver(semver string) bool {
	matched, err := regexp.Match(semverRegex, []byte(semver))
	if err != nil {
		panic(err)
	}
	return matched
}
