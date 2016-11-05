package lib

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/mayflower/docker-ls/lib/connector"
)

type registryApi struct {
	cfg       Config
	connector connector.Connector
}

func (r *registryApi) GetRegistryUrl() *url.URL {
	return &r.cfg.registryUrl
}

func (r *registryApi) endpointUrl(path string) *url.URL {
	url := r.cfg.registryUrl

	url.Path = path

	return &url
}

func (r *registryApi) paginatedRequestEndpointUrl(path string, lastApiResponse *http.Response) (url *url.URL, err error) {
	url = r.endpointUrl(path)

	if lastApiResponse != nil {
		linkHeader := lastApiResponse.Header.Get("link")

		if linkHeader != "" {
			// This is a hack to work around what looks like a bug in the registry:
			// the supplied link URL currently lacks scheme and host
			scheme, host := url.Scheme, url.Host

			url, err = parseLinkToNextHeader(linkHeader)

			if err != nil {
				return
			}

			if url.Scheme == "" {
				url.Scheme = scheme
			}

			if url.Host == "" {
				url.Host = host
			}
		}
	} else {
		queryParams := url.Query()
		queryParams.Set("n", strconv.Itoa(int(r.pageSize())))
		url.RawQuery = queryParams.Encode()
	}

	return
}

func (r *registryApi) pageSize() uint {
	return r.cfg.pageSize
}

func (r *registryApi) GetStatistics() connector.Statistics {
	return r.connector.GetStatistics()
}

func NewRegistryApi(cfg Config) (api RegistryApi, err error) {
	err = cfg.Validate()
	if err != nil {
		return
	}

	registry := &registryApi{
		cfg: cfg,
	}

	registry.connector = createConnector(&registry.cfg)

	api = registry
	return
}
