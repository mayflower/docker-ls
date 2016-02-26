package connector

import (
	"git.mayflower.de/vaillant-team/docker-ls/lib/auth"
)

type Config interface {
	MaxConcurrentRequests() uint
	Credentials() auth.RegistryCredentials
	AllowInsecure() bool
}
