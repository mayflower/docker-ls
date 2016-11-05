package lib

import (
	"errors"
	"flag"
	"net/url"
	"os"

	"github.com/mayflower/docker-ls/lib/auth"
)

var DEFAULT_REGISTRY_URL url.URL
var DEFAULT_REGISTRY_URL_STRING string

func init() {
	initRegistryURL()
}

func initRegistryURL() {
	DEFAULT_REGISTRY_URL_STRING = os.Getenv("DOCKER_REGISTRY_URL")
	if DEFAULT_REGISTRY_URL_STRING == "" {
		DEFAULT_REGISTRY_URL_STRING = "https://index.docker.io"
	}
	parsed, _ := url.Parse(DEFAULT_REGISTRY_URL_STRING)

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
	allowInsecure         bool
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
	flags.BoolVar(&c.allowInsecure, "allow-insecure", c.allowInsecure, "ignore SSL certificate validation errors")

	c.credentials.BindToFlags(flags)
}

func (c *Config) MaxConcurrentRequests() uint {
	return c.maxConcurrentRequests
}

func (c *Config) Credentials() auth.RegistryCredentials {
	return &c.credentials
}

func (c *Config) AllowInsecure() bool {
	return c.allowInsecure
}

func (c *Config) SetUrl(url url.URL) {
	c.registryUrl = url
}

func (c *Config) SetCredentials(credentials RegistryCredentials) {
	c.credentials = credentials
}

func (c *Config) SetPagesize(pageSize uint) {
	c.pageSize = pageSize
}

func (c *Config) SetMaxConcurrentRequests(maxRequests uint) {
	c.maxConcurrentRequests = maxRequests
}

func (c *Config) SetUseBasicAuth(basicAuth bool) {
	c.basicAuth = basicAuth
}

func (c *Config) SetAllowInsecure(allowInsecure bool) {
	c.allowInsecure = allowInsecure
}

func (c *Config) Validate() error {
	if c.pageSize == 0 {
		return errors.New("pagesize must be nonzero")
	}

	if c.maxConcurrentRequests == 0 {
		return errors.New("max requests must be nonzero")
	}

	return nil
}

func NewConfig() Config {
	return Config{
		registryUrl:           DEFAULT_REGISTRY_URL,
		pageSize:              100,
		maxConcurrentRequests: 5,
		basicAuth:             false,
	}
}
