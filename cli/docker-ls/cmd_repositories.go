package main

import (
	"flag"
	"fmt"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type repositoriesCmd struct {
	flags *flag.FlagSet
}

func (r *repositoriesCmd) execute(argv []string) (err error) {
	cfg := lib.NewConfig()
	cfg.BindToFlags(r.flags)

	err = r.flags.Parse(argv)

	if err != nil {
		return
	}

	registryApi := lib.NewRegistryApi(cfg)

	listResult, err := registryApi.ListRepositories()

	if err != nil {
		return
	} else {
		for repository := range listResult.Repositories() {
			fmt.Println(repository.Name())
		}

		err = listResult.LastError()
	}

	return
}

func newRepositoriesCmd(name string) (cmd *repositoriesCmd) {
	cmd = &repositoriesCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "", cmd.flags)

	return
}
