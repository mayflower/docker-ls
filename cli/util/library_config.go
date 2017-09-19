package util

import (
	"net/url"

	"github.com/mayflower/docker-ls/lib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var defaultConfig lib.Config = lib.NewConfig()

type LibraryFlags struct {
	RegistryUrl           string
	Pagesize              uint
	MaxConcurrentRequests uint
	BasicAuth             bool
	Username              string
	Password              string
	AllowInsecure         bool
}

func AddLibraryConfigToFlags(flags *pflag.FlagSet) {
	flags.StringP("registry", "r", defaultConfig.RegistryUrl().String(),
		"registry URL",
	)
	flags.Uint("page-size", defaultConfig.PageSize(),
		"page size for paginated requests",
	)
	flags.Uint("max-requests", defaultConfig.MaxConcurrentRequests(),
		"max. number of concurrent API requests",
	)
	flags.Bool("basic-auth", defaultConfig.UseBasicAuth(),
		"use basic auth instead of token-based auth",
	)
	flags.StringP("user", "u", defaultConfig.Credentials().User(),
		"username for registry login",
	)
	flags.StringP("password", "p", defaultConfig.Credentials().Password(),
		"password for registry login",
	)
	flags.Bool("allow-insecure", defaultConfig.AllowInsecure(),
		"ignore SSL certificate validation errors",
	)
}

func LibraryConfigFromViper() (config *lib.Config, err error) {
	c := lib.NewConfig()

	var parsedUrl *url.URL
	if parsedUrl, err = url.Parse(viper.GetString("registry")); err != nil {
		return
	}

	c.SetUrl(*parsedUrl)
	c.SetPagesize(uint(viper.GetInt("page-size")))
	c.SetMaxConcurrentRequests(uint(viper.GetInt("max-requests")))
	c.SetUseBasicAuth(viper.GetBool("basic-auth"))
	c.Credentials().SetPassword(viper.GetString("password"))
	c.Credentials().SetUser(viper.GetString("user"))
	c.SetAllowInsecure(viper.GetBool("allow-insecure"))

	config = &c
	return
}
