package lib

import (
	"flag"
)

type RegistryCredentials struct {
	user     string
	password string
}

func (c *RegistryCredentials) BindToFlags(flags *flag.FlagSet) {
	flags.StringVar(&c.user, "user", "", "username for logging into the registry")
	flags.StringVar(&c.password, "password", "", "password for logging into the registry")
}

func (r *RegistryCredentials) User() string {
	return r.user
}

func (r *RegistryCredentials) Password() string {
	return r.password
}

func (r *RegistryCredentials) SetPassword(password string) {
	r.password = password
}

func NewRegistryCredentials(user, password string) RegistryCredentials {
	return RegistryCredentials{
		user:     user,
		password: password,
	}
}
