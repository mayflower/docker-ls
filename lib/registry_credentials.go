package lib

import (
	"flag"
	"net/url"
	"os"

	"github.com/docker/cli/cli/config"
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

func (r *RegistryCredentials) SetUser(user string) {
	r.user = user
}

func (r *RegistryCredentials) Password() string {
	return r.password
}

func (r *RegistryCredentials) SetPassword(password string) {
	r.password = password
}

func (r *RegistryCredentials) IsBlank() bool {
	return r.User() == "" && r.Password() == ""
}

func (r *RegistryCredentials) LoadCredentialsFromDockerConfig(url url.URL) {
	authConfig, err := config.LoadDefaultConfigFile(os.Stderr).GetCredentialsStore(url.Host).Get(url.Host)

	if err != nil {
		return
	}

	r.SetUser(authConfig.Username)
	r.SetPassword(authConfig.Password)
}

func NewRegistryCredentials(user, password string) RegistryCredentials {
	return RegistryCredentials{
		user:     user,
		password: password,
	}
}
