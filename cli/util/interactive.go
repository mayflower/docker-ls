package util

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/mayflower/docker-ls/lib"
)

func PromptPassword(config *lib.Config) (err error) {
	credentials := config.Credentials()

	fmt.Fprintf(os.Stderr, "please enter password for user %s: ", credentials.User())

	binaryPassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return
	}

	credentials.SetPassword(string(binaryPassword))
	fmt.Fprintln(os.Stderr)

	return
}
