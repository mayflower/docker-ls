package main

import (
	"flag"
)

type tagsCmd struct {
	flags *flag.FlagSet
}

func (r *tagsCmd) execute(argv []string) (err error) {
	err = r.flags.Parse(argv)

	return
}

func newtagsCmd(name string) (cmd *tagsCmd) {
	cmd = &tagsCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "<respository>", cmd.flags)

	return
}
