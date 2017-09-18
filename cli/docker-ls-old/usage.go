package main

import (
	"flag"
	"fmt"
)

const COMMAND_USAGE_TEMPLATE = `usage: docker-ls %s [options] %s

%s

valid options:

`

func commandUsage(command, argstring, description string, flags *flag.FlagSet) func() {
	return func() {
		fmt.Printf(COMMAND_USAGE_TEMPLATE, command, argstring, description)
		flags.PrintDefaults()
	}
}
