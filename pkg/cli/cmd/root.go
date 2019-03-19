package cmd

import (
	"fmt"
	"log"
	"os"

	kitlog "github.com/go-kit/kit/log"
	"github.com/gobuffalo/packr"
	"github.com/meinto/glow/pkg/cli/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmdOptions struct {
	Author                string
	GitPath               string
	UseBuiltInGitBindings bool
	CICDOrigin            string
	CI                    bool
}

var logger kitlog.Logger

var rootCmd = &cobra.Command{
	Use:   "glow",
	Short: "small tool to adapt git-flow for gitlab",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if rootCmdOptions.CICDOrigin != "" {
			g, err := util.GetGitClient()
			util.CheckForError(err, "GetGitClient")

			err = g.SetCICDOrigin(rootCmdOptions.CICDOrigin)
			util.CheckForError(err, "SetCICDOrigin")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		g, err := util.GetGitClient()
		util.CheckForError(err, "GetGitClient")

		repoPath, err := g.GitRepoPath()
		util.CheckForError(err, "GitRepoPath")

		box := packr.NewBox(repoPath + "/buildAssets")
		version, err := box.FindString("VERSION")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Version of glow: %s\n", version)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootCmdOptions.Author, "author", "a", "name of the author")
	rootCmd.PersistentFlags().StringVar(&rootCmdOptions.GitPath, "gitPath", "/usr/local/bin/git", "path to native git installation")
	rootCmd.PersistentFlags().BoolVar(&rootCmdOptions.UseBuiltInGitBindings, "useBuiltInGitBindings", false, "defines wether build or native in git client should be used.")
	rootCmd.PersistentFlags().StringVar(&rootCmdOptions.CICDOrigin, "cicdOrigin", "", "provide a git origin url where a pipeline can push things via token")
	rootCmd.PersistentFlags().BoolVar(&rootCmdOptions.CI, "ci", false, "detects if command is running in a ci")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("gitPath", rootCmd.PersistentFlags().Lookup("gitPath"))
	viper.BindPFlag("useBuiltInGitBindings", rootCmd.PersistentFlags().Lookup("useBuiltInGitBindings"))
}

func Execute() {
	logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
