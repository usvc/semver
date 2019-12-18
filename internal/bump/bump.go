package bump

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
	Name:      "bump",
	ShortName: "+",
	Usage:     "bump a provided semver",
	Description: `
	 bump a cli literal input:
		 semver bump v1.2.3

	 bump a cli piped input:
		 echo 'v1.2.3' | semver bump

	 bump the latest git tag:
		 semver bump --git

	 bump the latest git tag and add it:
		 semver bump --git --apply`,
	ArgsUsage:    "<semver version>",
	HideHelp:     false,
	BashComplete: BashComplete,
	Flags:        Flags,
	Action:       Action,
}

var Flags = []cli.Flag{
	cli.BoolFlag{
		Name:  "major, M",
		Usage: "bump the major version",
	},
	cli.BoolFlag{
		Name:  "minor, m",
		Usage: "bump the minor version",
	},
	cli.BoolFlag{
		Name:  "patch, p",
		Usage: "bump the patch version (defaults to this)",
	},
	cli.BoolFlag{
		Name:  "label, l",
		Usage: "bump the label iteration",
	},
	cli.BoolFlag{
		Name:  "git",
		Usage: "use latest version retrieved from git tags",
	},
	cli.BoolFlag{
		Name:  "apply",
		Usage: "if specified with --git, adds the bumped version automatically to the tags",
	},
	cli.BoolFlag{
		Name:  "verbose, vv",
		Usage: "show verbose logs (do not enable for automations)",
	},
}

func Action(c *cli.Context) error {
	input := "v0.0.0"
	if c.Bool("git") {
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
		if len(tags) > 0 {
			input = tags[len(tags)-1].ToString()
		}
	} else {
		input = utils.GetCLIInput(c)
	}

	if len(input) == 0 {
		cli.ShowSubcommandHelp(c)
		os.Exit(-1)
		return nil
	}

	if !semver.IsValid(input) {
		err := fmt.Errorf("'%s' does not seem to be a semver version", input)
		fmt.Println(err)
		os.Exit(1)
		return err
	}

	semverInput := semver.New(input)
	currentVersion := semverInput.ToString()
	switch {
	case c.Bool("label"):
		semverInput.BumpLabel()
	case c.Bool("patch"):
		semverInput.BumpPatch()
	case c.Bool("minor"):
		semverInput.BumpMinor()
	case c.Bool("major"):
		semverInput.BumpMajor()
	default:
		semverInput.BumpPatch()
	}
	if c.Bool("git") && c.Bool("apply") {
		git.AddTag(utils.GetCurrentWorkingDirectory(), semverInput.ToString())
		fmt.Printf("added git tag '%s'", semverInput.ToString())
	} else {
		if c.Bool("verbose") {
			// separated so that only the version is output in non-verbose mode
			fmt.Printf("%s -> ", currentVersion)
		}
		fmt.Print(semverInput.ToString())
	}
	if c.Bool("verbose") {
		fmt.Printf("\n")
	}
	os.Exit(0)
	return nil
}

func BashComplete(c *cli.Context) {
	if c.NArg() > 0 {
		return
	}
	fmt.Println("--major")
	fmt.Println("--minor")
	fmt.Println("--patch")
	fmt.Println("--label")
	fmt.Println("--git")
	fmt.Println("--apply")
	fmt.Println("--verbose")
}
