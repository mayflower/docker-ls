package connector

import (
	"git.mayflower.de/vaillant-team/docker-ls/lib/auth"
)

type TokenAuthConfig interface {
	MaxConcurrentRequests() uint
	Credentials() auth.RegistryCredentials
}
