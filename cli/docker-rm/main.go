package main

import (
	"flag"
	"fmt"
	"os"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

const USAGE_TEMPLATE = `usage: docker-rm <repository> <reference> [options]

Delete a tag in a given repository.

valid options:

`

var flags *flag.FlagSet = flag.NewFlagSet("main", flag.ExitOnError)

func init() {
	flags.Usage = usage
}

func usage() {
	fmt.Printf(USAGE_TEMPLATE)

	flags.PrintDefaults()
}

func main() {
	libCfg := lib.NewConfig()
	libCfg.BindToFlags(flags)

	if len(os.Args) < 3 {
		usage()
		os.Exit(0)
	}

	repository, reference := os.Args[1], os.Args[2]

	flags.Parse(os.Args[3:])

	api := lib.NewRegistryApi(libCfg)

	if err := api.DeleteTag(repository, reference); err == nil {
		fmt.Println("...Tag deleted successfully!")
	} else {
		fmt.Printf("ERROR: %v\n", err)
	}
}
