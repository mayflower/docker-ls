package lib

import (
	"flag"
	"net/url"
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
	registryUrl url.URL
	credentials RegistryCredentials
	pageSize    uint
}

func (u *urlValue) String() string {
	return (*url.URL)(u).String()
}

func (c *Config) BindToFlags(flags *flag.FlagSet) {
	c.registryUrl = DEFAULT_REGISTRY_URL

	flags.Var((*urlValue)(&c.registryUrl), "registry", "registry URL")
	flags.UintVar(&c.pageSize, "page-size", 100, "page size for paginated requests")
	c.credentials.BindToFlags(flags)
}

func NewConfig() Config {
	return Config{
		registryUrl: DEFAULT_REGISTRY_URL,
	}
}
