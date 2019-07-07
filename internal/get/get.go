package get

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
	"gitlab.com/usvc/utils/semver"
	"gitlab.com/usvc/utils/semver/internal/utils"
	"gitlab.com/usvc/utils/semver/pkg/git"
)

var Command = cli.Command{
	Name:      "get",
	ShortName: "-",
	Usage:     "gets the current semver",
	Description: `
	 retrieve the current git tag:
		 semver get`,
	HideHelp:     false,
	Flags:        Flags,
	Action:       Action,
}

var Flags = []cli.Flag{
	cli.StringFlag{
		Name: "use",
		Usage: "defines the engine to use (defaults to 'git')",
	},
}

func Action(c *cli.Context) error {
	switch {
	case c.String("use") == "git":
		fallthrough
	default:
		var tags semver.Semvers
		cwd := utils.GetCurrentWorkingDirectory()
		retrievedTags, retrieveTagsErr := git.GetTags(cwd)
		if retrieveTagsErr != nil {
			err := fmt.Errorf("an error happened while retrieving git tags from '%s': %s", cwd, retrieveTagsErr)
			fmt.Println(err)
			os.Exit(1)
		}
		for _, t := range retrievedTags {
			if semver.IsValid(t) {
				tags = append(tags, semver.New(t))
			}
		}
		sort.Sort(tags)
		fmt.Println(tags[len(tags)-1].ToString())
	}
	return nil
}

func BashComplete(c *cli.Context) {
	if c.NArg() > 0 {
		return
	}
	fmt.Println("--use")
}