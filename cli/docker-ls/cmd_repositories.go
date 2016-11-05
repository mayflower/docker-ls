package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/mayflower/docker-ls/cli/docker-ls/response"
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
)

type repositoriesCmd struct {
	flags *flag.FlagSet
	cfg   *Config
}

func (r *repositoriesCmd) execute(argv []string) (err error) {
	libCfg := lib.NewConfig()
	libCfg.BindToFlags(r.flags)

	r.cfg = newConfig()
	r.cfg.bindToFlags(r.flags,
		OPTION_JSON_OUTPUT|OPTION_PROGRESS|OPTION_RECURSION_LEVEL|OPTION_STATISTICS|OPTION_INTERACTIVE_PASSWORD|OPTION_TABLE_OUTPUT)

	err = r.flags.Parse(argv)
	if err != nil {
		return
	}

	if r.cfg.interactivePassword {
		err = util.PromptPassword(&libCfg)
		if err != nil {
			return
		}
	}

	if len(r.flags.Args()) != 0 {
		r.flags.Usage()
		os.Exit(1)
	}

	registryApi, err := lib.NewRegistryApi(libCfg)
	if err != nil {
		return
	}

	var resp sortable

	switch {
	case r.cfg.recursionLevel >= 1:
		resp, err = r.listLevel1(registryApi)

	case r.cfg.recursionLevel == 0:
		resp, err = r.listLevel0(registryApi)
	}

	if err != nil {
		return
	}

	resp.Sort()

	if r.cfg.tableOutput {
		w := tabwriter.NewWriter(os.Stdout, 50, 1, 3, ' ', 0)
		switch repositories := resp.(type) {
		case *response.RepositoriesL0:
			fmt.Fprintln(w, "REPOSITORY")
			for _, repository := range repositories.Repositories {
				fmt.Fprintf(w, "%s\n", repository)
			}
			w.Flush()
		case *response.RepositoriesL1:
			fmt.Fprintln(w, "REPOSITORY\tTAG")
			for _, repository := range repositories.Repositories {
				for _, tag := range repository.Tags {
					fmt.Fprintf(w, "%s\t%s\n", repository.RepositoryName, tag)
				}
			}
			w.Flush()
		}
	} else {
		err = serializeToStdout(resp, r.cfg)
	}

	if r.cfg.statistics {
		dumpStatistics(registryApi.GetStatistics())
	}

	return
}

func (r *repositoriesCmd) listLevel0(api lib.RegistryApi) (resp *response.RepositoriesL0, err error) {
	progress := NewProgressIndicator(r.cfg)
	progress.Start(fmt.Sprintf("requesting list from %s", api.GetRegistryUrl()))

	result := api.ListRepositories()
	resp = response.NewRepositoriesL0()

	progress.Progress()

	for repository := range result.Repositories() {
		resp.AddRepository(repository)
	}

	err = result.LastError()

	progress.Finish("done")
	return
}

func (r *repositoriesCmd) listLevel1(api lib.RegistryApi) (resp *response.RepositoriesL1, err error) {
	progress := NewProgressIndicator(r.cfg)
	progress.Start(fmt.Sprintf("requesting list from %s", api.GetRegistryUrl()))

	repositoriesResult := api.ListRepositories()
	resp = response.NewRepositoriesL1()
	progress.Progress()

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

				progress.Progress()
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

	progress.Finish("done")
	return
}

func newRepositoriesCmd(name string) (cmd *repositoriesCmd) {
	cmd = &repositoriesCmd{
		flags: flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.flags.Usage = commandUsage(name, "", "List all repositories.", cmd.flags)

	return
}
