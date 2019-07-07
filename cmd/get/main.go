//go:generate go run ../../scripts/versioning/main.go

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	Get "gitlab.com/usvc/utils/semver/internal/get"
)

func main() {
	app := cli.NewApp()
	app.Name = "semver-get"
	app.Version = fmt.Sprintf("%s [%s]", Version, Commit)
	app.Author = "@zephinzer <dev-at-joeir-dot-net>"
	app.Description = "get semver versions"
	app.Usage = "get semver semvers"
	app.EnableBashCompletion = true
	app.BashComplete = Get.BashComplete
	app.Flags = Get.Flags
	app.Action = Get.Action

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
