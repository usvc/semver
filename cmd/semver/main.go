package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "semver"
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
