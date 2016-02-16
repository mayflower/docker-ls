package main

import (
	"flag"
	"os"

	"git.mayflower.de/vaillant-team/docker-ls/cli/docker-ls/response"
	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type tagsCmd struct {
	flags *flag.FlagSet
}

func (r *tagsCmd) execute(argv []string) (err error) {
	cfg := lib.NewConfig()
	cfg.BindToFlags(r.flags)

	if len(argv) == 0 {
		r.flags.Usage()
		os.Exit(1)
	}

	repositoryName := argv[0]

	err = r.flags.Parse(argv[1:])

	if err != nil {
		return
	}

	registryApi := lib.NewRegistryApi(cfg)

	listResult := registryApi.ListTags(repositoryName)
	resp := response.NewTagsL0(repositoryName)

	for tag := range listResult.Tags() {
		resp.AddTag(tag)
	}

	err = listResult.LastError()
	if err != nil {
		return
	}

	err = yamlToStdout(resp)

	return
}

func newTagsCmd(name string) (cmd *tagsCmd) {
	cmd = &tagsCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "<respository>", cmd.flags)

	return
}
