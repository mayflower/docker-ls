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

func (f *LibraryFlags) BindToFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&f.RegistryUrl, "registry", "r", defaultConfig.RegistryUrl().String(),
		"registry URL",
	)
	flags.UintVar(&f.Pagesize, "page-size", defaultConfig.PageSize(),
		"page size for paginated requests",
	)
	flags.UintVar(&f.MaxConcurrentRequests, "max-requests", defaultConfig.MaxConcurrentRequests(),
		"max. number of concurrent API requests",
	)
	flags.BoolVar(&f.BasicAuth, "basic-auth", defaultConfig.UseBasicAuth(),
		"use basic auth instead of token-based auth",
	)
	flags.StringVarP(&f.Username, "user", "u", defaultConfig.Credentials().User(),
		"username for registry login",
	)
	flags.StringVarP(&f.Password, "password", "p", defaultConfig.Credentials().Password(),
		"password for registry login",
	)
	flags.BoolVar(&f.AllowInsecure, "allow-insecure", defaultConfig.AllowInsecure(),
		"ignore SSL certificate validation errors",
	)

	viper.BindPFlags(flags)
}

func (f *LibraryFlags) CreateLibraryConfig() (config *lib.Config, err error) {
	c := lib.NewConfig()

	var url *url.URL
	if url, err = url.Parse(viper.GetString("registry")); err != nil {
		return
	}

	c.SetUrl(*url)
	c.SetPagesize(uint(viper.GetInt("page-size")))
	c.SetMaxConcurrentRequests(uint(viper.GetInt("max-requests")))
	c.SetUseBasicAuth(viper.GetBool("basic-auth"))
	c.Credentials().SetPassword(viper.GetString("password"))
	c.Credentials().SetUser(viper.GetString("user"))
	c.SetAllowInsecure(viper.GetBool("allow-insecure"))

	config = &c
	return
}
