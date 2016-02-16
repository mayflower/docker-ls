package main

import (
	"flag"
	"os"

	"git.mayflower.de/vaillant-team/docker-ls/cli/docker-ls/response"
	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type tagDetailsCmd struct {
	flags *flag.FlagSet
}

func (r *tagDetailsCmd) execute(argv []string) (err error) {
	cfg := lib.NewConfig()
	cfg.BindToFlags(r.flags)

	if len(argv) < 2 {
		r.flags.Usage()
		os.Exit(1)
	}

	repositoryName := argv[0]
	reference := argv[1]

	err = r.flags.Parse(argv[2:])

	if err != nil {
		return
	}

	registryApi := lib.NewRegistryApi(cfg)
	tagDetails, err := registryApi.GetTagDetails(repositoryName, reference)

	if err != nil {
		return
	}

	err = yamlToStdout(response.NewTagDetailsL0(tagDetails, true))

	return
}

func newTagDetailsCmd(name string) (cmd *tagDetailsCmd) {
	cmd = &tagDetailsCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "<respository> <reference>", cmd.flags)

	return
}
