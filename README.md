[![Build Status](https://travis-ci.org/mayflower/docker-ls.svg?branch=master)](https://travis-ci.org/mayflower/docker-ls)

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

Six ways there are to attain enlightenment.

## Precompiled binaries

Just download precompiled binaries for your platform from
[github](https://github.com/mayflower/docker-ls/releases).

## MacOS / Homebrew

You can install `docker-ls` directly from homebrew:

    brew install docker-ls

## Gentoo / portage

```
emerge docker-ls
```

## NixOS

```
nix-env -iA nixos.docker-ls
```

## Arch Linux

Package in the [AUR](https://aur.archlinux.org/packages/docker-ls/) available.


## Docker

If you have Docker installed, you may want to try this option. Clone the
repository and do:

    docker build -t docker-ls .

Example of running container:

    $ docker run -it docker-ls docker-ls tags library/consul
    requesting list . done
    repository: library/consul
    tags:
    - latest
    - v0.6.4

Or create aliases:

    $ alias docker-ls='docker run -it docker-ls docker-ls'
    $ alias docker-rm='docker run -it docker-ls docker-rm'

So you can do:

    $ docker-ls tags library/consul
    requesting list . done
    repository: library/consul
    tags:
    - latest
    - v0.6.4

and:

    $ docker-rm | head -n 3
    usage: docker-rm [options] <repository:reference>

    Delete a tag in a given repository.

## Go get

Provided that you sport an installation of
[golang](https://golang.org), the latest version from master
can be installed via

    go get -d github.com/mayflower/docker-ls/cli/...
    go generate github.com/mayflower/docker-ls/lib/...
    go install github.com/mayflower/docker-ls/cli/...

Isn't a simple `go get github.com/mayflower/docker-ls/cli/...` sufficient, you ask?
Indeed it is, but including the generate step detailed above will encode verbose version information
in the binaries.

# Usage

Docker-ls contains two CLI tools: `docker-ls` and `docker-rm` .

## docker-ls

`docker-ls` is a browser for docker registries. Output is either encoded as YAML or
as JSON.

Several subcommands are available

 * `docker-ls repositories` Obtains a list of repositories on the server.
   **This is not supported by the official [docker hub](https://hub.docker.com/).**
 * `docker-ls tags` Lists all tags in a a particular repository.
 * `docker-ls tag` Inspect a particular tag. This command displays a condensed version
   of the corresponding manifest by default, but the `--raw-manifest` option can be
   used to dump the full manifest. The `--parse-history` option can be used to display
   the JSON-encoded history within the manifest.

### Authentication and credentials

By default, `docker-ls` uses the token based authentication flow for authentication unless
basic auth is requested explicitly (see below). If no crendentials are specified, `docker-ls`
will automatically get the credentials from the docker CLI (if logged in via `docker login`).
`docker-ls` implicitly uses the same credential helpers userd by docker.

Logging into Amazon ECR requires basic auth.

### Important command line flags

This list is not exhaustive; please consult the command line (`-h`) help for all options.

 * `--registry <url> (-r)` Connect to the registry at <url>. The URL must include the protocol
   (http / https). By default, `docker-ls` targets the official
   [docker hub](https://hub.docker.com/).
 * `--user <user> (-u)` Username for authentication.
 * `--password <password> (-p)` Password for authentication.
 * `--user-agent <agent string>` Use a custom user agent.
 * `--interactive-password(-i)` Read the password from an interactive prompt.
 * `--level <depth> (-l)` The `repositories` and `tags` subcommands support this option
   for recursive output. Depths 0 (default) and 1 are supported. Please note
   the recursing means more API requests and may be slow.
 * `--json (-j)` Switch output format from YAML to JSON.
 * `--template (-t)` Use a named golang template from the configuration for output (see below)
 * `--template-source` Use the specified template for output (see below)
 * `--basic-auth` Use HTTP basic auth for authentication (instead of token authentication).
 * `--allow-insecure` Do not validate SSL certificates (useful for registries secured with a
    self-signed certificate).
 * `--manifest-version` Request either manifest version
   [V2.1](https://github.com/docker/distribution/blob/master/docs/spec/manifest-v2-1.md)
   (`--manifest-version 1` or manifest version [V2.2](https://github.com/docker/distribution/blob/master/docs/spec/manifest-v2-2.md)
   (`--manifest-version 2`, default) from the registry. Please note that deleting manifests
   from registry version >= 2.3 will work **only** with content digests from a V2.2
   manifest.

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

### Notes considering the offical registry

If no registry is specified, `docker-ls` will target the official registry server
at `https://index.docker.io`. Please note that:

 * The official registry does not support repository listing via `docker-ls repositories`
 * Official repositories must be prefixed with `library/`, e.g. `docker-ls tags library/debian`

## docker-rm

`docker-rm` can delete particular tags. Example:

    docker-rm --registry https://my.registry.org --user someuser --password somepass busybox:sha256:51fef[...]

(the digest has been truncated for brevity). Please consult the command line help
for a full list of all arguments.

Some remarks:

 * The tag *must* be specified as a sha256 content digest.
 * While tags can be deleted, the current registry implementation will (to the best
   of my knowledge) not free the space associated with any resulting unused layers
 * Deleting stuff is currently disabled by default in the official registry and needs to be
   enabled explicitly &mdash; check out this [issue](https://github.com/mayflower/docker-ls/issues/1)
   for details.
 * Content digests obtained with `--manifest-version 1` will **not work** with
   registry version >= 2.3.
 * **BE CAREFUL!** The API does not implement undelete :)

## Using a proxy

`docker-ls` supports HTTP / HTTPS proxies configured via the corresponding
canonical environment variables. Check out the corresponding
[documentation](https://golang.org/pkg/net/http/#ProxyFromEnvironment)
for details.

## Configuration via config files and environment variables

All options that can be specified via CLI flags can be read from a config file or from an
environment variables. The priority is CLI flag > environment variable > config file.

### Config files

By default, both tools try to read
`~/.docker-ls.[yaml|json|toml|...]`
(please check the Viper [documentation](https://github.com/spf13/viper)
for a full list of the supported formats). The names of the keys in the file
are the long names of the CLI flags. For example, the following YAML file would configure
registry URL and username

    registry: https://foo.bar.com
    user: foo

Other config files can be specified via the `--config` option.

### Template Output

Output of the various `docker-ls` subcommands can be further customized by using
[golang templates](https://golang.org/pkg/text/template/).

#### Predefined templates

Named templates can be configured in the `templates` section of the configuration file.
When `docker-ls` is invoked, the `-t` parameter (see above) can be used to select a named
template for formatting the output.

**Example:** The following YAML section defines a template that outputs the list of tags
in a repository as a simple HTML document.

```
templates:
  taglist_html: |
    <head></head>
    <body>
        <h1>Tags for repository {{ html .Repository }}</h1>
        <ul>
            {{- range .Tags }}
            <li>{{ html . }}</li>
            {{- end }}
        </ul>
    </body>
```

It can be invoked by running i.e.

```
docker-ls tags -t taglist_html /library/debian
```

#### Inline templates

Simple templates can also be passed directly on the command line using the `--template-source`
parameter:

```
docker-ls tag --template-source '{{ .TagName }}: {{ .Digest }}'  /library/debian:wheezy
```

### Template variables

Inside templates, all fields of the corresponding JSON / YAML output can be accessed in pipeline
expressions. The first letter of all field names is capitalized, with the exception of manifests
that are directly returned from the registry by using `docker-ls tag --raw-manifest`: for
those, the JSON / YAML field names are unchanged.

### Environment variables

In addition to config files and CLI flags, environment variables can be used to specify options
globally. The name is determined by taking the long CLI name, uppercasing replacing
hyphens "-" with underscores "\_" and prefixing the result with "DOCKER_LS_". For example,
the following would enable interactive password prompts for all consecutive
invocations:

    export DOCKER_LS_INTERACTIVE_PASSWORD=1

## Shell autocompletion

Both `docker-ls` and `docker-rm` support shell autocompletion for subcommands and
options. To enable this, source the output of `docker-ls autocomplete bash|zsh`
and `docker-rm autocomplete bash|zsh` in the running shell. In case of bash, this can be
achieved with

    $ source <(docker-ls autocomplete bash)

# License

Docker-ls is distributed under the terms of the MIT license.
