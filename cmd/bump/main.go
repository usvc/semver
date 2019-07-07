//go:generate go run ../../scripts/versioning/main.go

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	Bump "gitlab.com/usvc/utils/semver/internal/bump"
)

func main() {
	app := cli.NewApp()
	app.Name = "semver-bump"
	app.Version = fmt.Sprintf("%s [%s]", Version, Commit)
	app.Author = "@zephinzer <dev-at-joeir-dot-net>"
	app.Description = "bump semver versions"
	app.Usage = "bump semver semvers"
	app.EnableBashCompletion = true
	app.BashComplete = Bump.BashComplete
	app.Flags = Bump.Flags
	app.Action = Bump.Action

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
