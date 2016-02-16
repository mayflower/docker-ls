package main

import (
	"flag"
	"fmt"
)

const COMMAND_USAGE_TEMPLATE = `usage: docker-ls %s%s [options]

valid options:

`

func commandUsage(command string, argstring string, flags *flag.FlagSet) func() {
	return func() {
		fmt.Printf(COMMAND_USAGE_TEMPLATE, command, argstring)
		flags.PrintDefaults()
	}
}
