package main

import (
	"flag"
	"fmt"
	"os"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

const USAGE_TEMPLATE = `usage: docker-rm <repository:reference> [options]

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

func dispatch() (err error) {
	libCfg := lib.NewConfig()
	libCfg.BindToFlags(flags)

	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	ref := lib.EmptyRefspec()
	err = ref.Set(os.Args[1])
	if err != nil {
		return
	}

	flags.Parse(os.Args[2:])

	api := lib.NewRegistryApi(libCfg)
	err = api.DeleteTag(ref)

	return
}

func main() {
	if err := dispatch(); err == nil {
		fmt.Println("...Tag deleted successfully!")
	} else {
		fmt.Printf("ERROR: %v\n", err)
	}
}
