package tags

import (
	"sync"

	"github.com/mayflower/docker-ls/cli/docker-ls/response"
	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
)

type Executor struct {
	CliConfig     *util.CliConfig
	LibraryConfig *lib.Config
	Repository    string
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
	case e.CliConfig.RecursionLevel == 0:
		resp, err = e.listLevel0(registryApi)

	case e.CliConfig.RecursionLevel >= 1:
		resp, err = e.listLevel1(registryApi)
	}

	if err != nil {
		return
	}

	resp.Sort()

	if e.CliConfig.Template != "" {
		err = util.TemplateToStdout(resp, e.CliConfig.Template)
	} else {
		err = util.SerializeToStdout(resp, e.CliConfig)
	}

	if e.CliConfig.Statistics {
		util.DumpStatistics(registryApi.GetStatistics())
	}

	return
}

func (e *Executor) listLevel0(api lib.RegistryApi) (resp *response.TagsL0, err error) {
	progress := util.NewProgressIndicator(e.CliConfig)
	progress.Start("requesting list")

	listResult := api.ListTags(e.Repository)
	progress.Progress()
	resp = response.NewTagsL0(e.Repository)

	for tag := range listResult.Tags() {
		resp.AddTag(tag)
	}

	err = listResult.LastError()

	progress.Finish("done")
	return
}

func (e *Executor) listLevel1(api lib.RegistryApi) (resp *response.TagsL1, err error) {
	progress := util.NewProgressIndicator(e.CliConfig)
	progress.Start("requesting list")

	listResult := api.ListTags(e.Repository)
	progress.Progress()
	resp = response.NewTagsL1(e.Repository)

	errors := make(chan error)

	go func() {
		var wait sync.WaitGroup

		for tag := range listResult.Tags() {
			wait.Add(1)

			go func(tag lib.Tag) {
				tagDetails, err := api.GetTagDetails(lib.NewRefspec(tag.RepositoryName(), tag.Name()), e.CliConfig.ManifestVersion)
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
