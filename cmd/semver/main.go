package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/usvc/utils/semver/cmd/semver/bump"
)

var (
	Version   string
	Commit    string
	Timestamp string
)

func run(command *cobra.Command, args []string) {
	command.Help()
}

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "semver",
		Version: fmt.Sprintf("%s-%s %s", Version, Commit, Timestamp),
		Run:     run,
	}
	cmd.AddCommand(bump.GetCommand())
	return cmd
}

func main() {
	GetCommand().Execute()
}
