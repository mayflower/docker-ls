package lib

import (
	"errors"
	"flag"
	"log"
	"net/url"

	"github.com/mayflower/docker-ls/lib/auth"
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
	allowInsecure         bool
	userAgent             string
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
	flags.StringVar(&c.userAgent, "user-agent", c.userAgent, "override http user-agent header")

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

func (c *Config) UserAgent() string {
	return c.userAgent
}

func (c *Config) RegistryUrl() *url.URL {
	return &c.registryUrl
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

func (c *Config) PageSize() uint {
	return c.pageSize
}

func (c *Config) SetMaxConcurrentRequests(maxRequests uint) {
	c.maxConcurrentRequests = maxRequests
}

func (c *Config) SetUseBasicAuth(basicAuth bool) {
	c.basicAuth = basicAuth
}

func (c *Config) UseBasicAuth() bool {
	return c.basicAuth
}

func (c *Config) SetAllowInsecure(allowInsecure bool) {
	c.allowInsecure = allowInsecure
}

func (c *Config) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
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

func (c *Config) LoadCredentialsFromDockerConfig() {
	log.Printf("load credentials? %t", c.credentials.IsBlank())
	if c.credentials.IsBlank() {
		c.credentials.LoadCredentialsFromDockerConfig(c.registryUrl)
	}
}

func NewConfig() Config {
	return Config{
		registryUrl:           DEFAULT_REGISTRY_URL,
		pageSize:              100,
		maxConcurrentRequests: 5,
		basicAuth:             false,
		userAgent:             ApplicationName(),
	}
}
