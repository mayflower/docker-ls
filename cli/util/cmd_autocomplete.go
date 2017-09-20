package util

import (
	"os"

	"github.com/spf13/cobra"
)

var autocompleteBashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Autocompletion snippet for bash",
	Long:  "Generate autocompletion snippet for bash",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().GenBashCompletion(os.Stdout)
	},
}

var autocompleteZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Autocompletion snippet for zsh",
	Long:  "Generate autocompletion snippet for zsh",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().GenZshCompletion(os.Stdout)
	},
}

var AutocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generate autocompletion snippet",
	Long:  "Generate autocompletion snippet for either bash or zsh",
}

func init() {
	AutocompleteCmd.AddCommand(autocompleteBashCmd, autocompleteZshCmd)
}
