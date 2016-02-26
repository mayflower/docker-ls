package connector

import (
	"net/http"
	"net/url"
)

type Connector interface {
	Request(method string, url *url.URL, hint string) (*http.Response, error)
	Delete(url *url.URL, hint string) (*http.Response, error)
	Get(url *url.URL, hint string) (*http.Response, error)
	GetStatistics() Statistics
}
