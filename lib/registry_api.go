package lib

import (
	"net/url"
)

type RegistryApi struct {
	cfg       Config
	connector *RegistryConnector
	pageSize  int
}

func (r *RegistryApi) endpointUrl(path string) *url.URL {
	url := r.cfg.registryUrl

	url.Path = path

	return &url
}

func NewRegistryApi(cfg Config) (registry *RegistryApi) {
	registry = &RegistryApi{
		cfg:       cfg,
		connector: NewRegistryConnector(cfg),
		pageSize:  20,
	}

	return
}
