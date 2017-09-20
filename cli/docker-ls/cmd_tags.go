package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/mayflower/docker-ls/cli/docker-ls/tags"
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
	"github.com/spf13/cobra"
)

var tagsCommand = &cobra.Command{
	Use:   "tags",
	Short: "List tags",
	Long:  "List all tags for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())

		var err error
		var libraryConfig *lib.Config
		if libraryConfig, err = util.LibraryConfigFromViper(); err == nil {
			executor := tags.Executor{
				CliConfig:     util.CliConfigFromViper(),
				LibraryConfig: libraryConfig,
				Repository:    args[0],
			}

			err = executor.Execute()
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(tagsCommand)

	flags := tagsCommand.Flags()

	util.AddCliConfigToFlags(flags, util.CLI_OPTIONS_FULL)
	util.AddLibraryConfigToFlags(flags)
}
