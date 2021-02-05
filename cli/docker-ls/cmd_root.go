package main

import (
	"log"
	"os"

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

		if debug {
			log.SetOutput(os.Stdout)
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Printf("debug logging enabled")
		}
	},
}

func init() {
	var configFile string
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"read config from specified file (default: look for config in home directory)",
	)
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "emit debugging logs")

	util.SetupViper(configFile)
}
