package minor

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/usvc/go-logger"
	"github.com/usvc/go-semver"
	bumpUtils "gitlab.com/usvc/utils/semver/cmd/semver/bump/utils"
	"gitlab.com/usvc/utils/semver/cmd/semver/common"
	"gitlab.com/usvc/utils/semver/internal/utils"
)

var (
	cmd      *cobra.Command
	errorLog = logger.New(logger.Options{
		Type:   logger.TypeStdout,
		Output: logger.OutputStderr,
	})
	log = logger.New(logger.Options{
		Type: logger.TypeStdout,
	})
)

func GetCommand() *cobra.Command {
	if cmd == nil {
		cmd = &cobra.Command{
			Use: "minor <semver>",
			Run: func(command *cobra.Command, args []string) {
				hasArguments := (len(args) > 0)
				retrieveFromGit := common.Config.GetBool("git")
				applyToGit := common.Config.GetBool("apply")
				var version *semver.Semver
				var err error
				switch true {
				case retrieveFromGit:
					version, err = bumpUtils.GetCurrentRepositoryLatestSemver()
				case hasArguments:
					version, err = utils.GetSemverFromArguments(args)
				default:
					command.Help()
					return
				}
				if err != nil {
					errorLog.Errorf("failed to retrieve latest version: '%s'", err)
					os.Exit(1)
				}
				version.BumpMinor()
				log.Info(version.String())
				if retrieveFromGit && applyToGit {
					log.Debugf("adding git tag '%s' to repository...", version.String())
					err = bumpUtils.SetCurrentRepositorySemver(version)
					log.Infof("added tag '%s'", version.String)
				}
			},
		}
		common.Config.ApplyToCobra(cmd)
	}
	return cmd
}
