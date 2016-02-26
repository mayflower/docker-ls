# What is it?

Docker-ls is a set of CLI tools for browsing and manipulating docker registries.
In particular, docker-ls can handle authentication and display the sha256 content
digests associated with tags.

## What registries are supported

Only
[V2 registries](https://github.com/docker/distribution)
 are supported. Both HTTP basic auth and docker style
[token authentication](https://github.com/docker/distribution/blob/master/docs/spec/auth/token.md)
are supported for authentication.

# Installation

Three ways there are to attain enlightenment.

## Precompiled binaries

Just download precompiled binaries for your platform from
[github](https://github.com/mayflower/docker-ls/releases).

## Go get

Provided that you sport an installation of
[golang](git@github.com:mayflower/docker-ls.git), the latest version from master
can be installed via

    go get -d github.com/mayflower/docker-ls/...
    go generate github.com/mayflower/docker-ls/...
    go install github.com/mayflower/docker-ls/cli/...

Isn't a simple `go get github.com/mayflower/docker-ls/cli/...` sufficient, you ask?
Indeed it is, but including the generate step detailed above will encode verbose version information
in the binaries.

## Git & Make

Clone the repository and do `make`. This will create a separate `GOPATH` in `build`
and leave you with the binaries ready in `build/bin`. Of course, you need
[golang](git@github.com:mayflower/docker-ls.git)
installed for this.

# Usage

Docker-ls contains two CLI tools: `docker-ls` and `docker-rm` .

## docker-ls

`docker-ls` is a browser for docker registries. Output is either encoded as YAML or
as JSON.

Three subcommands are available

 * `docker-ls repositories` Obtains a list of repositories on the server.
   **This is not supported by the official docker hub.**
 * `docker-ls tags` Lists all tags in a a particular repository.
 * `docker-ls tag` Inspect a particular tag. This command displays a condensed version
   of the corresponding manifest by default, but the `--raw-manifest` option can be
   used to dump the full manifest. The `--parse-history` option can be used to display
   the JSON-encoded history within the manifest.

### Important command line flags

This list is not exhaustive; please consult the command line (`-h`) help for all options.

 * `--registry <url>` Connect to the registry at <url>. The URL must include the protocol
   (http / https).
 * `--user <user>` Username for authentication.
 * `--password <password>` Password for authentication.
 * `--level <depth>` The `repositories` and `tags` subcommands support this option
   for recursive output. Depths 0 (default) and 1 are supported. Please note
   the recursing means more API requests and may be slow.
 * `--json` Switch output format from YAML to JSON.
 * `--basic-auth` Use HTTP basic auth for authentication (instead of token authentication).
 * `--allow-insecure` Do not validate SSL certificates (useful for registries secured with a
    self-signed certificate).

### Examples

List all repositories in a custom registry:

    docker-ls repositories --registry https://my.registry.org --user hanni --password hanni123

List all repositories in a custom registry, including their tags:

    docker-ls repositories --registry https://my.registry.org --user hanni --password hanni123 --level 1

List all tags in stuff/busybox using HTTP basic auth

    docker-ls tags --registry https://my.registry.org --user hanni --password hanni123 --basic-auth stuff/busybox

Inspect tag stuff/busybox:latest, no authentication, JSON output.

    docker-ls tag --registry https://my.registry.org --json stuff/busybox:latest

Inspect tag stuff/busybox:latest, no authentication, dump the raw manifest with parsed
history as JSON.

    docker-ls tag --registry https://my.registry.org --json --raw-manifest --parse-history stuff/busybox:latest

## docker-rm

`docker-rm` can delete particular tags. Example:

    docker-rm --registry https://my.registry.org --user someuser --password somepass busybox/sha256:51fef[...]

(the digest has been truncated for brevity). Please consult the command line help
for a full list of all arguments.

Some remarks:

 * The tag *must* be specified as a sha256 content digest.
 * While tags can be deleted, the current registry implementation will (to the best
   of my knowledge) not free the space associated with any resulting unused layers.
 * **BE CAREFUL!** The API does not implement undelete :)

# License

`Docker-ls` is distributed under the terms of the MIT license.
