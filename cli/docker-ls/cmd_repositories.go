package main

import (
	"flag"
	"sync"

	"git.mayflower.de/vaillant-team/docker-ls/cli/docker-ls/response"
	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type repositoriesCmd struct {
	flags *flag.FlagSet
}

func (r *repositoriesCmd) execute(argv []string) (err error) {
	libCfg := lib.NewConfig()
	libCfg.BindToFlags(r.flags)

	cfg := newConfig()
	cfg.bindToFlags(r.flags)

	err = r.flags.Parse(argv)

	if err != nil {
		return
	}

	registryApi := lib.NewRegistryApi(libCfg)
	var resp sortable

	switch {
	case cfg.recursionLevel >= 1:
		resp, err = r.listLevel1(registryApi)

	case cfg.recursionLevel == 0:
		resp, err = r.listLevel0(registryApi)
	}

	if err != nil {
		return
	}

	resp.Sort()
	err = yamlToStdout(resp)

	return
}

func (r *repositoriesCmd) listLevel0(api lib.RegistryApi) (resp *response.RepositoriesL0, err error) {
	result := api.ListRepositories()
	resp = response.NewRepositoriesL0()

	for repository := range result.Repositories() {
		resp.AddRepository(repository)
	}

	err = result.LastError()

	return
}

func (r *repositoriesCmd) listLevel1(api lib.RegistryApi) (resp *response.RepositoriesL1, err error) {
	repositoriesResult := api.ListRepositories()
	resp = response.NewRepositoriesL1()

	errors := make(chan error)

	go func() {
		var wait sync.WaitGroup

		for repository := range repositoriesResult.Repositories() {
			wait.Add(1)

			go func(repository lib.Repository) {
				tagsResult := api.ListTags(repository.Name())
				tagsL0 := response.NewTagsL0(repository.Name())

				for tag := range tagsResult.Tags() {
					tagsL0.AddTag(tag)
				}

				resp.AddTags(tagsL0)

				if err := tagsResult.LastError(); err != nil {
					errors <- err
				}

				wait.Done()
			}(repository)
		}

		if err := repositoriesResult.LastError(); err != nil {
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

func newRepositoriesCmd(name string) (cmd *repositoriesCmd) {
	cmd = &repositoriesCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "", cmd.flags)

	return
}
