package major

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/usvc/go-logger"
	"github.com/usvc/go-semver"
	"gitlab.com/usvc/utils/semver/cmd/semver/common"
)

var (
	cmd *cobra.Command
	log logger.Logger
)

func run(command *cobra.Command, args []string) {
	hasArguments := (len(args) > 0)
	retrieveFromGit := common.Config.GetBool("git")
	var version *semver.Semver
	var err error

	switch true {
	case retrieveFromGit:
		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Errorf("failed to get current working directory: '%s'", err)
			os.Exit(1)
		}
		version, err = common.GetLatestSemverFromGitRepository(currentDirectory)
		if err != nil {
			log.Errorf("failed to retrieve latest semver tag: '%s'", err)
			os.Exit(1)
		}
	case hasArguments:
		version, err = common.GetSemverFromArguments(args)
		if err != nil {
			log.Errorf("failed to parse semver input from arguments: '%s'", err)
			os.Exit(1)
		}
	default:
		command.Help()
		return
	}

	log.Debugf("current version: '%s'", version)
	version.BumpPatch()
	log.Debugf("next version: '%s'", version)

	if common.Config.GetBool("git") && common.Config.GetBool("apply") {
		log.Debugf("adding git tag '%s' to repository...", version.String())

		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Errorf("failed to getting current working directory: '%s'", err)
			os.Exit(1)
		}
		err = common.TagCurrentGitCommit(version.String(), currentDirectory)
		if err != nil {
			log.Errorf("failed to add tag '%s' to repository at '%s': '%s'", version.String, currentDirectory, err)
			os.Exit(1)
		}
		log.Infof("added tag '%s' to repository at '%s'", version.String, currentDirectory)
	}
}

func GetCommand() *cobra.Command {
	if cmd == nil {
		log = logger.New(logger.Options{})
		common.Config.LoadFromEnvironment()
		cmd = &cobra.Command{
			Use: "major",
			Run: run,
		}
		common.Config.ApplyToCobra(cmd)
	}
	return cmd
}
