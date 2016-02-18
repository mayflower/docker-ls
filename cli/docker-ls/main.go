package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Print(`usage: docker-ls command [options]

valid commands:

    repositories    List all repositories
    tags            List all tags for a single repository
    tag             Fetch details for an individual tag
    version         Show version
`)

	os.Exit(1)
}

func parseCommandLine() string {
	if len(os.Args) <= 1 {
		usage()
	}

	return os.Args[1]
}

func getCommand() command {
	switch parseCommandLine() {
	case "repositories":
		return newRepositoriesCmd("respositories")

	case "tags":
		return newTagsCmd("tags")

	case "tag":
		return newTagDetailsCmd("tag")

	case "version":
		return newVersionCmd()

	default:
		return nil
	}

}

func main() {
	command := getCommand()

	if command == nil {
		usage()
	}

	err := command.execute(os.Args[2:])

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
