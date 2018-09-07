package main

import (
	"fmt"
	"os"

	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "docker-rm <repository:tag>",
	Short: "Delete a tag",
	Long:  "Delete a tag in a given repository",
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
			executor := Executor{
				CliConfig:     cliConfig,
				LibraryConfig: libraryConfig,
				Tag:           args[0],
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
	var configFile string
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"read config from specified file (default: look for config in home directory)",
	)

	flags := rootCmd.Flags()
	util.AddCliConfigToFlags(flags, util.CLI_OPTION_INTERACTIVE_PASSWORD)
	util.AddLibraryConfigToFlags(flags)

	util.SetupViper(configFile)
}
