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
	libCfg := lib.NewConfig()
	libCfg.BindToFlags(r.flags)

	cfg := newConfig()
	cfg.bindToFlags(r.flags, OPTION_PROGRESS|OPTION_JSON_OUTPUT)

	rawManifest := false
	r.flags.BoolVar(&rawManifest, "raw-manifest", rawManifest, "output raw manifest")

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

	progress := NewProgressIndicator(cfg)
	progress.Start("requesting manifest")

	registryApi := lib.NewRegistryApi(libCfg)
	tagDetails, err := registryApi.GetTagDetails(repositoryName, reference)

	progress.Progress()
	progress.Finish("done")

	if err != nil {
		return
	}

	if rawManifest {
		err = serializeToStdout(tagDetails.RawManifest(), cfg)
	} else {
		err = serializeToStdout(response.NewTagDetailsL0(tagDetails, true), cfg)
	}

	return
}

func newTagDetailsCmd(name string) (cmd *tagDetailsCmd) {
	cmd = &tagDetailsCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "<respository> <reference>", "Inspect a singe tag.", cmd.flags)

	return
}
