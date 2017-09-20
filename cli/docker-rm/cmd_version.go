package main

import "github.com/mayflower/docker-ls/cli/util"

func init() {
	rootCmd.AddCommand(util.VersionCmd)
}
