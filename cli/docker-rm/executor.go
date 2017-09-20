package main

import (
	"fmt"

	"github.com/mayflower/docker-ls/cli/util"
	"github.com/mayflower/docker-ls/lib"
)

type Executor struct {
	CliConfig     *util.CliConfig
	LibraryConfig *lib.Config
	Tag           string
}

func (e *Executor) Execute() (err error) {
	if e.CliConfig.InteractivePassword {
		err = util.PromptPassword(e.LibraryConfig)
		if err != nil {
			return
		}
	}

	ref := lib.EmptyRefspec()
	err = ref.Set(e.Tag)
	if err != nil {
		return
	}

	api, err := lib.NewRegistryApi(*e.LibraryConfig)
	if err != nil {
		return
	}

	if err = api.DeleteTag(ref); err == nil {
		fmt.Println("...Tag deleted successfully!")
	}

	return
}
