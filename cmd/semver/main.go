//go:generate go run ../../scripts/versioning/main.go

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "semver"
	app.Version = fmt.Sprintf("%s [%s]", Version, Commit)
	app.Author = "@zephinzer <dev-at-joeir-dot-net>"
	app.Description = "easy-peasy manipulation of semver versions"
	app.Usage = "easy-peasy manipulation of semver versions"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		bump,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
