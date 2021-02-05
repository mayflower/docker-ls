package main

import (
	"fmt"
	"os"

	"github.com/mayflower/docker-ls/cli/docker-ls/tags"
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags <repository>",
	Short: "List tags",
	Long:  "List all tags for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		var libraryConfig *lib.Config
		libraryConfig, err = util.LibraryConfigFromViper()

		var cliConfig *util.CliConfig
		if err == nil {
			cliConfig, err = util.CliConfigFromViper()
		}

		if err == nil {
			executor := tags.Executor{
				CliConfig:     cliConfig,
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
	rootCmd.AddCommand(tagsCmd)

	flags := tagsCmd.Flags()

	util.AddCliConfigToFlags(
		flags,
		util.CLI_OPTIONS_FULL & ^util.CLI_OPTION_TABLE_OUTPUT,
	)
	util.AddLibraryConfigToFlags(flags)
}
