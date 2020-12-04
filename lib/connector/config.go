package connector

import (
	"github.com/mayflower/docker-ls/lib/auth"
)

type Config interface {
	MaxConcurrentRequests() uint
	Credentials() auth.RegistryCredentials
	AllowInsecure() bool
	UserAgent() string
}
