package main

import (
	"flag"
	"os"
	"sync"

	"github.com/mayflower/docker-ls/cli/docker-ls/response"
	"github.com/mayflower/docker-ls/lib"
)

type tagsCmd struct {
	flags          *flag.FlagSet
	repositoryName string
	cfg            *Config
}

func (r *tagsCmd) execute(argv []string) (err error) {
	libCfg := lib.NewConfig()
	libCfg.BindToFlags(r.flags)

	r.cfg = newConfig()
	r.cfg.bindToFlags(r.flags, OPTIONS_FULL)

	err = r.flags.Parse(argv)
	if err != nil {
		return
	}

	args := r.flags.Args()
	if len(args) != 1 {
		r.flags.Usage()
		os.Exit(1)
	}
	r.repositoryName = args[0]

	registryApi, err := lib.NewRegistryApi(libCfg)
	if err != nil {
		return
	}

	var resp sortable

	switch {
	case r.cfg.recursionLevel == 0:
		resp, err = r.listLevel0(registryApi)

	case r.cfg.recursionLevel >= 1:
		resp, err = r.listLevel1(registryApi)
	}

	if err != nil {
		return
	}

	resp.Sort()

	err = serializeToStdout(resp, r.cfg)

	if r.cfg.statistics {
		dumpStatistics(registryApi.GetStatistics())
	}

	return
}

func (r *tagsCmd) listLevel0(api lib.RegistryApi) (resp *response.TagsL0, err error) {
	progress := NewProgressIndicator(r.cfg)
	progress.Start("requesting list")

	listResult := api.ListTags(r.repositoryName)
	progress.Progress()
	resp = response.NewTagsL0(r.repositoryName)

	for tag := range listResult.Tags() {
		resp.AddTag(tag)
	}

	err = listResult.LastError()

	progress.Finish("done")
	return
}

func (r *tagsCmd) listLevel1(api lib.RegistryApi) (resp *response.TagsL1, err error) {
	progress := NewProgressIndicator(r.cfg)
	progress.Start("requesting list")

	listResult := api.ListTags(r.repositoryName)
	progress.Progress()
	resp = response.NewTagsL1(r.repositoryName)

	errors := make(chan error)

	go func() {
		var wait sync.WaitGroup

		for tag := range listResult.Tags() {
			wait.Add(1)

			go func(tag lib.Tag) {
				tagDetails, err := api.GetTagDetails(lib.NewRefspec(tag.RepositoryName(), tag.Name()))
				progress.Progress()

				if err == nil {
					resp.AddTag(tagDetails)
				} else {
					errors <- err
				}

				wait.Done()
			}(tag)
		}

		if err := listResult.LastError(); err != nil {
			errors <- err
		}

		wait.Wait()

		close(errors)
	}()

	for nextError := range errors {
		if err == nil {
			err = nextError
		}
	}

	progress.Finish("done")
	return
}

func newTagsCmd(name string) (cmd *tagsCmd) {
	cmd = &tagsCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "<respository>", "Show all tags for a given repository.", cmd.flags)

	return
}
