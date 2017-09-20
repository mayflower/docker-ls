package util

import (
	"fmt"

	"github.com/mayflower/docker-ls/lib"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show docker-ls version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", lib.Version())
	},
}
