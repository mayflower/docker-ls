package main

import (
	"flag"
	"fmt"
	"strings"
)

const COMMAND_USAGE_TEMPLATE = `usage: docker-ls %s%s [options]

%s

valid options:

`

func commandUsage(command, argstring, description string, flags *flag.FlagSet) func() {
	return func() {
		if argstring != "" {
			argstring = " " + strings.TrimLeft(argstring, " ")
		}

		fmt.Printf(COMMAND_USAGE_TEMPLATE, command, argstring, description)
		flags.PrintDefaults()
	}
}
