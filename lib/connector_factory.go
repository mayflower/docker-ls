package lib

import (
	"github.com/mayflower/docker-ls/lib/connector"
)

func createConnector(cfg *Config) connector.Connector {
	if cfg.basicAuth {
		return connector.NewBasicAuthConnector(cfg)
	}
	return connector.NewTokenAuthConnector(cfg)
}
