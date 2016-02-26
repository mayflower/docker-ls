package lib

import (
	"flag"
	"net/url"

	"git.mayflower.de/vaillant-team/docker-ls/lib/auth"
)

var DEFAULT_REGISTRY_URL url.URL

func init() {
	parsed, _ := url.Parse("https://index.docker.io")

	DEFAULT_REGISTRY_URL = *parsed
}

type urlValue url.URL

func (u *urlValue) Set(value string) (err error) {
	var parsedUrl *url.URL
	parsedUrl, err = url.Parse(value)

	if err == nil {
		*(*url.URL)(u) = *parsedUrl
	}

	return
}

type Config struct {
	registryUrl           url.URL
	credentials           RegistryCredentials
	pageSize              uint
	maxConcurrentRequests uint
	basicAuth             bool
}

func (u *urlValue) String() string {
	return (*url.URL)(u).String()
}

func (c *Config) BindToFlags(flags *flag.FlagSet) {
	c.registryUrl = DEFAULT_REGISTRY_URL

	flags.Var((*urlValue)(&c.registryUrl), "registry", "registry URL")
	flags.UintVar(&c.pageSize, "page-size", c.pageSize, "page size for paginated requests")
	flags.UintVar(&c.maxConcurrentRequests, "max-requests", c.maxConcurrentRequests, "concurrent API request limit")
	flags.BoolVar(&c.basicAuth, "basic-auth", c.basicAuth, "use basic auth instead of token auth")

	c.credentials.BindToFlags(flags)
}

func (c *Config) MaxConcurrentRequests() uint {
	return c.maxConcurrentRequests
}

func (c *Config) Credentials() auth.RegistryCredentials {
	return &c.credentials
}

func NewConfig() Config {
	return Config{
		registryUrl:           DEFAULT_REGISTRY_URL,
		pageSize:              100,
		maxConcurrentRequests: 5,
		basicAuth:             false,
	}
}
