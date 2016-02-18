package main

import (
	"fmt"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type versionCmd struct{}

func (v versionCmd) execute(argv []string) error {
	fmt.Printf("version: %s\n", lib.Version())

	return nil
}

func newVersionCmd() versionCmd {
	return versionCmd{}
}
