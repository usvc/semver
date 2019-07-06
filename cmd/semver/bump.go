package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
	"gitlab.com/usvc/utils/semver"
	"gitlab.com/usvc/utils/semver/pkg/git"
)

var bump cli.Command = cli.Command{
	Name:      "bump",
	ShortName: "+",
	Usage:     "bumps a provided semver",
	ArgsUsage: "[ eg. \"v1.2.3\" ]",

	BashComplete: func(c *cli.Context) {
		if c.NArg() > 0 {
			return
		}
		fmt.Println("--major")
		fmt.Println("--minor")
		fmt.Println("--patch")
	},
	Flags: bumpFlags,
	Action: func(c *cli.Context) error {
		var input string
		if c.Bool("git") {
			var tags semver.Semvers
			cwd := getCurrentWorkingDirectory()
			retrievedTags, retrieveTagsErr := git.GetTags(cwd)
			if retrieveTagsErr != nil {
				err := fmt.Errorf("an error happened while retrieving git tags from '%s'", cwd)
				fmt.Println(err)
				os.Exit(1)
			}
			for _, t := range retrievedTags {
				if semver.IsValid(t) {
					tags = append(tags, semver.New(t))
				}
			}
			sort.Sort(tags)
			input = tags[len(tags)-1].ToString()
		} else {
			input = getCLIInput(c)
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
		if c.Bool("verbose") {
			fmt.Printf("%s -> ", semverInput.ToString())
		}
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
		fmt.Printf(semverInput.ToString())
		if c.Bool("verbose") {
			fmt.Printf("\n")
		}
		os.Exit(0)
		return nil
	},
}

var bumpFlags []cli.Flag = []cli.Flag{
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
		Name:  "verbose, vv",
		Usage: "show verbose logs (do not enable for automations)",
	},
	cli.BoolFlag{
		Name:  "git",
		Usage: "use latest version retrieved from git tags",
	},
}
