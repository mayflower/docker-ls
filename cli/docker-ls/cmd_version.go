package main

import (
	"fmt"

	"github.com/mayflower/go-repro/lib"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show docker-ls version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", lib.Version())
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
