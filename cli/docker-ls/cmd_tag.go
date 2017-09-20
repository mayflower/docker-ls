package main

import (
	"fmt"
	"os"

	"github.com/mayflower/docker-ls/cli/docker-ls/tag"
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tagCmd = &cobra.Command{
	Use:   "tag <repository:tag>",
	Short: "Show tag details",
	Long:  "Detailed inspection of a particular tag",
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())

		var err error
		var libraryConfig *lib.Config
		if libraryConfig, err = util.LibraryConfigFromViper(); err == nil {
			executor := tag.Executor{
				CliConfig:     util.CliConfigFromViper(),
				LibraryConfig: libraryConfig,
				Tag:           args[0],
				RawManifest:   viper.GetBool("raw-manifest"),
				ParseHistory:  viper.GetBool("parse-history"),
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
	rootCmd.AddCommand(tagCmd)

	flags := tagCmd.Flags()

	util.AddCliConfigToFlags(
		flags,
		util.CLI_OPTION_PROGRESS|
			util.CLI_OPTION_JSON_OUTPUT|
			util.CLI_OPTION_MANIFEST_VERSION|
			util.CLI_OPTION_INTERACTIVE_PASSWORD,
	)
	util.AddLibraryConfigToFlags(flags)

	flags.Bool("raw-manifest", false, "dump raw manifest")
	flags.Bool("parse-history", false, "attempt to parse history in raw manifest")
}
