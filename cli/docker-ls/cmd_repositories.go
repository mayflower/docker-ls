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

		var cliConfig *util.CliConfig
		if err == nil {
			cliConfig, err = util.CliConfigFromViper()
		}

		if err == nil {
			executor := repositories.Executor{
				CliConfig:     cliConfig,
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
		util.CLI_OPTIONS_FULL & ^util.CLI_OPTION_MANIFEST_VERSION,
	)
}
