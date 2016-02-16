package main

import (
	"flag"

	"git.mayflower.de/vaillant-team/docker-ls/cli/docker-ls/response"
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

	listResult := registryApi.ListRepositories()
	resp := response.NewRepositoriesL0()

	for repository := range listResult.Repositories() {
		resp.AddRepository(repository)
	}

	err = listResult.LastError()
	if err != nil {
		return
	}

	err = yamlToStdout(resp)

	return
}

func newRepositoriesCmd(name string) (cmd *repositoriesCmd) {
	cmd = &repositoriesCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "", cmd.flags)

	return
}
