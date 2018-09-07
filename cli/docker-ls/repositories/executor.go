package repositories

import (
	"fmt"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/mayflower/docker-ls/cli/docker-ls/response"
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
)

type Executor struct {
	CliConfig     *util.CliConfig
	LibraryConfig *lib.Config
}

func (e *Executor) Execute() (err error) {
	if e.CliConfig.InteractivePassword {
		err = util.PromptPassword(e.LibraryConfig)
		if err != nil {
			return
		}
	}

	registryApi, err := lib.NewRegistryApi(*e.LibraryConfig)
	if err != nil {
		return
	}

	var resp util.Sortable

	switch {
	case e.CliConfig.RecursionLevel >= 1:
		resp, err = e.listLevel1(registryApi)

	case e.CliConfig.RecursionLevel == 0:
		resp, err = e.listLevel0(registryApi)
	}

	if err != nil {
		return
	}

	resp.Sort()

	if e.CliConfig.TableOutput {
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
					fmt.Fprintf(w, "%s\t%s\n", repository.Repository, tag)
				}
			}
			w.Flush()
		}
	} else {
		err = util.SerializeToStdout(resp, e.CliConfig)
	}

	if e.CliConfig.Statistics {
		util.DumpStatistics(registryApi.GetStatistics())
	}

	return
}

func (e *Executor) listLevel0(api lib.RegistryApi) (resp *response.RepositoriesL0, err error) {
	progress := util.NewProgressIndicator(e.CliConfig)
	progress.Start("requesting list")

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

func (e *Executor) listLevel1(api lib.RegistryApi) (resp *response.RepositoriesL1, err error) {
	progress := util.NewProgressIndicator(e.CliConfig)
	progress.Start("requesting list")

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
