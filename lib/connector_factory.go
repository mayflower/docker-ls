package lib

import (
	"git.mayflower.de/vaillant-team/docker-ls/lib/connector"
)

func createConnector(cfg *Config) connector.Connector {
	return connector.NewTokenAuthConnector(cfg)
}
