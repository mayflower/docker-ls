package main

import (
	"flag"
	"fmt"
	"os"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

const USAGE_TEMPLATE = `usage: [options] docker-rm <repository:reference>

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

func version() {
	fmt.Printf("version: %s\n", lib.Version())
}

func dispatch() (err error) {
	libCfg := lib.NewConfig()
	libCfg.BindToFlags(flags)

	showVersion := false
	flags.BoolVar(&showVersion, "version", false, "show version and exit")

	flags.Parse(os.Args[1:])

	if showVersion {
		version()
		os.Exit(0)
	}

	args := flags.Args()
	if len(args) != 1 {
		usage()
		os.Exit(1)
	}

	ref := lib.EmptyRefspec()
	err = ref.Set(args[0])
	if err != nil {
		return
	}

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
