package main

import (
	"flag"
	"fmt"
	"os"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

var flags *flag.FlagSet

func init() {
	flags = flag.NewFlagSet("default", flag.ExitOnError)
	flags.Usage = usage
}

func usage() {
	fmt.Print(`usage: docker-ls [options]

valid options:

`)

	flags.PrintDefaults()
}

func parseCommandLine() (cfg lib.Config) {
	cfg = lib.Config{}
	cfg.BindToFlags(flags)

	flags.Parse(os.Args[1:])

	return
}

func main() {
	cfg := parseCommandLine()

	registryApi := lib.NewRegistryApi(cfg)

	listResult, err := registryApi.ListRepositories()

	if err != nil {
		fmt.Println(err)
	} else {
		for repository := range listResult.Repositories() {
			fmt.Println(repository.Name())
		}

		if err := listResult.LastError(); err != nil {
			fmt.Printf("\nERROR: %v\n", err)
		}
	}
}
