package bump

import (
	"github.com/spf13/cobra"
	"github.com/usvc/go-logger"
	"gitlab.com/usvc/utils/semver/cmd/semver/bump/major"
	"gitlab.com/usvc/utils/semver/cmd/semver/bump/minor"
	"gitlab.com/usvc/utils/semver/cmd/semver/bump/patch"
	"gitlab.com/usvc/utils/semver/cmd/semver/bump/prerelease"
	"gitlab.com/usvc/utils/semver/cmd/semver/common"
)

var (
	cmd *cobra.Command
	log logger.Logger
)

func GetCommand() *cobra.Command {
	if cmd == nil {
		log = logger.New(logger.Options{})
		cmd = &cobra.Command{
			Use: "bump <semver>",
			Run: func(command *cobra.Command, args []string) {
				command.Run(patch.GetCommand(), args)
			},
		}
		common.Config.ApplyToCobra(cmd)
		cmd.AddCommand(major.GetCommand())
		cmd.AddCommand(minor.GetCommand())
		cmd.AddCommand(patch.GetCommand())
		cmd.AddCommand(prerelease.GetCommand())
	}
	return cmd
}
