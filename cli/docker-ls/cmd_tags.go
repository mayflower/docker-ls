package main

import (
	"flag"
	"fmt"
	"os"

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

	for tag := range listResult.Tags() {
		fmt.Println(tag.Name())
	}

	err = listResult.LastError()

	return
}

func newtagsCmd(name string) (cmd *tagsCmd) {
	cmd = &tagsCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "<respository>", cmd.flags)

	return
}
