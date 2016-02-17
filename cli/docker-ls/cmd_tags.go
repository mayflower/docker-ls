package main

import (
	"flag"
	"os"
	"sync"

	"git.mayflower.de/vaillant-team/docker-ls/cli/docker-ls/response"
	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type tagsCmd struct {
	flags          *flag.FlagSet
	repositoryName string
}

func (r *tagsCmd) execute(argv []string) (err error) {
	libCfg := lib.NewConfig()
	libCfg.BindToFlags(r.flags)

	cfg := newConfig()
	cfg.bindToFlags(r.flags)

	if len(argv) == 0 {
		r.flags.Usage()
		os.Exit(1)
	}

	r.repositoryName = argv[0]

	err = r.flags.Parse(argv[1:])

	if err != nil {
		return
	}

	registryApi := lib.NewRegistryApi(libCfg)
	var resp sortable

	switch {
	case cfg.recursionLevel == 0:
		resp, err = r.listLevel0(registryApi)

	case cfg.recursionLevel >= 1:
		resp, err = r.listLevel1(registryApi)
	}

	if err != nil {
		return
	}

	resp.Sort()

	err = yamlToStdout(resp)

	if cfg.statistics {
		dumpStatistics(registryApi.GetStatistics())
	}

	return
}

func (r *tagsCmd) listLevel0(api lib.RegistryApi) (resp *response.TagsL0, err error) {
	listResult := api.ListTags(r.repositoryName)
	resp = response.NewTagsL0(r.repositoryName)

	for tag := range listResult.Tags() {
		resp.AddTag(tag)
	}

	err = listResult.LastError()

	return
}

func (r *tagsCmd) listLevel1(api lib.RegistryApi) (resp *response.TagsL1, err error) {
	listResult := api.ListTags(r.repositoryName)
	resp = response.NewTagsL1(r.repositoryName)

	errors := make(chan error)

	go func() {
		var wait sync.WaitGroup

		for tag := range listResult.Tags() {
			wait.Add(1)

			go func(tag lib.Tag) {
				tagDetails, err := api.GetTagDetails(tag.RepositoryName(), tag.Name())

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

	return
}

func newTagsCmd(name string) (cmd *tagsCmd) {
	cmd = &tagsCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "<respository>", cmd.flags)

	return
}
