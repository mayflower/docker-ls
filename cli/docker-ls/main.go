package main

import (
	"flag"
	"fmt"
	"io"
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

	connector := lib.NewRegistryConnector(cfg)

	response, err := connector.Get("v2/_catalog")

	fmt.Println(err)

	if err == nil {
		if response.Close {
			defer response.Body.Close()
		}

		io.Copy(os.Stdout, response.Body)

		fmt.Println()
	}
}
