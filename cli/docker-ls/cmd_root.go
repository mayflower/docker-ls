package main

import (
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-ls",
	Short: "browse docker registries",
	Long:  "Browse and examine repositories and tags in a docker registry",
}

func init() {
	var configFile string
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"read config from specified file (default: look for config in home directory)",
	)

	util.SetupViper(configFile)
}
