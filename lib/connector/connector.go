package connector

import (
	"net/http"
	"net/url"
)

type Connector interface {
	Request(method string, url *url.URL, headers map[string]string, hint string) (*http.Response, error)
	Delete(url *url.URL, headers map[string]string, hint string) (*http.Response, error)
	Get(url *url.URL, headers map[string]string, hint string) (*http.Response, error)
	GetStatistics() Statistics
}
