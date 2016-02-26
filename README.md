# What is it?

Docker-ls is a set of CLI tools for browsing and manipulating docker registries.

## What registries are supported

Only V2 registries are supported. Both HTTP basic auth and docker style
[token authentication](https://github.com/docker/distribution/blob/master/docs/spec/auth/token.md)
are supported for authentication.

# Installation

TODO

# Usage

Docker-ls contains two CLI tools: `docker-ls` and `docker-rm` . The following paragraphs
give an overview over their usage; please consult the CLI help (option `-h`) for
more details.

## docker-ls

`docker-ls` is capable of browsing docker registries. Three subcommands are available

 * `docker-ls repositories`: Obtains a list of repositories on the server.
   **This is not supported by the official docker hub.**
 * `docker-ls tags`: Lists all tags in a a particular repository.
 * `docker-ls tag`: Inspect a particular tag. This command displays a condensed version
   of the corresponding manifest by default, but the `--raw-manifest` option can be
   used to dump the full manifest. The `--parse-history` option can be used to display
   the JSON-encoded history within the manifest.

### Important command line flags

This list is not comprehensive; please consult the command line (`-h`) help for all options.

 * `--registry <url>` Connect to the registry at URL. The URL must include the protocol
   (http / https).
 * `--user <user>` Username for authentication.
 * `--password <password>` Password for authentication.
 * `--level <depth>` The `repositories` and `tags` subcommands support the level option
   for recursing into more details. Depth 0 (default) and 1 are supported. Please note
   the recursing means more requests and is slower.
 * `--json` Switch output format from YAML to JSON.
 * `--basic-auth` Use HTTP basic auth for authentication (instead of token authentication).
 * `--allow-insecure` Do not validate SSL certificates (useful for registries secured with a
    self-signed cert).

### Examples

List all repositories in a custom registry:

    docker-ls repositories --registry https://my.registry.org --user hanni --password hanni123

List all repositories in a custom registry, including their tags:

    docker-ls repositories --registry https://my.registry.org --user hanni --password hanni123 --level 1

List all tags in stuff/busybox using HTTP basic auth

    docker-ls tags --registry https://my.registry.org --user hanni --password hanni123 --basic-auth stuff/busybox

Inspect tag stuff/busybox:latest, no authentication, JSON outut.

    docker-ls tag --registry https://my.registry.org --json stuff/busybox:latest


Inspect tag stuff/busybox:latest, no authentication, dump the raw manifest with parsed
history as JSON.

    docker-ls tag --registry https://my.registry.org --json --raw-manifest --parse-history stuff/busybox:latest

## docker-rm

`docker-rm` can delete particular tags. Example:

    docker-rm --registry https://foo.bar.org --user someuser --password somepass busybox/sha256:51fef[...]

(the sha256 has been truncated for brevity). Please consult the command line help
for a full list of all arguments.

Some remarks:

 * The tag *must* be specified as a content sha256.
 * While tags can be deleted, the current registry implementation will (to the best
   of my knowledge) not free the space associated with any now-unused layers.
 * **BE CAREFUL!** The API does not implement undelete :)
