package cmd

import (
	"log"

	"github.com/meinto/glow"
	. "github.com/meinto/glow/pkg/cli/cmd/util"
	"github.com/meinto/glow/semver"
	"github.com/spf13/cobra"
)

func init() {
	closeCmd.AddCommand(closeReleaseCmd)
	AddFlagsForMergeRequests(closeReleaseCmd)
}

var closeReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "close a release branch",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]

		if version == "current" {
			g, err := GetGitClient()
			ExitOnError(err)

			pathToRepo, _, _, err := g.GitRepoPath()
			ExitOnError(err)

			s := semver.NewSemverService(
				pathToRepo,
				"/bin/bash",
				releaseCmdOptions.VersionFile,
				releaseCmdOptions.VersionFileType,
			)
			s = semver.NewLoggingService(s)
			v, err := s.GetCurrentVersion()
			ExitOnError(err)
			version = v
		}

		gp, err := GetGitProvider()
		ExitOnError(err)

		currentBranch, err := glow.NewRelease(version)
		ExitOnError(err)

		err = gp.Close(currentBranch)
		if !MergeRequestFlags.Gracefully {
			ExitOnError(err)
		} else {
			log.Println(err)
		}
	},
}
