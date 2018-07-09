package main

import (
	"fmt"
	"os"

	"github.com/mayflower/docker-ls/cli/docker-ls/repositories"
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var repositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "List repositories",
	Long:  "List all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())

		var err error

		var libraryConfig *lib.Config
		libraryConfig, err = util.LibraryConfigFromViper()

		if err == nil {
			executor := repositories.Executor{
				CliConfig:     util.CliConfigFromViper(),
				LibraryConfig: libraryConfig,
			}

			err = executor.Execute()
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
	Args: cobra.NoArgs,
}

func init() {
	flags := repositoriesCmd.Flags()

	rootCmd.AddCommand(repositoriesCmd)

	util.AddLibraryConfigToFlags(flags)
	util.AddCliConfigToFlags(
		flags,
		util.CLI_OPTION_JSON_OUTPUT|
			util.CLI_OPTION_PROGRESS|
			util.CLI_OPTION_RECURSION_LEVEL|
			util.CLI_OPTION_STATISTICS|
			util.CLI_OPTION_INTERACTIVE_PASSWORD|
			util.CLI_OPTION_TABLE_OUTPUT,
	)
}
