package main

import (
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:   "docker-ls",
	Short: "browse docker registries",
	Long:  "Browse and examine repositories and tags in a docker registry",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())
	},
}

func init() {
	util.SetupViper(rootCmd)

	util.AddCliConfigToFlags(rootCmd.PersistentFlags(), util.CLI_OPTION_DEBUG)
}
