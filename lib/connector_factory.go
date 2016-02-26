package lib

import (
	"git.mayflower.de/vaillant-team/docker-ls/lib/connector"
)

func createConnector(cfg *Config) connector.Connector {
	if cfg.basicAuth {
		return connector.NewBasicAuthConnector(cfg)
	}
	return connector.NewTokenAuthConnector(cfg)
}
