package main

import (
	"flag"
	"fmt"
	"strings"
)

const COMMAND_USAGE_TEMPLATE = `usage: docker-ls %s%s [options]

valid options:

`

func commandUsage(command string, argstring string, flags *flag.FlagSet) func() {
	return func() {
		if argstring != "" {
			argstring = " " + strings.TrimLeft(argstring, " ")
		}

		fmt.Printf(COMMAND_USAGE_TEMPLATE, command, argstring)
		flags.PrintDefaults()
	}
}
