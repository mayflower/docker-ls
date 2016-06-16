package connector

import (
	"net/http"
	"net/url"
	"strings"
)

type basicAuthConnector struct {
	cfg        Config
	httpClient *http.Client
	semaphore  semaphore
	stat       *statistics
}

func (r *basicAuthConnector) Delete(url *url.URL, headers map[string]string, hint string) (*http.Response, error) {
	return r.Request("DELETE", url, headers, hint)
}

func (r *basicAuthConnector) Get(url *url.URL, headers map[string]string, hint string) (*http.Response, error) {
	return r.Request("GET", url, headers, hint)
}

func (r *basicAuthConnector) GetStatistics() Statistics {
	return r.stat
}

func (r *basicAuthConnector) Request(
	method string,
	url *url.URL,
	headers map[string]string,
	hint string,
) (response *http.Response, err error) {
	r.semaphore.Lock()
	defer r.semaphore.Unlock()

	r.stat.Request()

	request, err := http.NewRequest(method, url.String(), strings.NewReader(""))

	if err != nil {
		return
	}

	credentials := r.cfg.Credentials()
	if credentials.Password() != "" || credentials.User() != "" {
		request.SetBasicAuth(credentials.User(), credentials.Password())
	}

	for header, value := range headers {
		request.Header.Set(header, value)
	}

	response, err = r.httpClient.Do(request)

	return
}

func NewBasicAuthConnector(cfg Config) Connector {
	return &basicAuthConnector{
		cfg:        cfg,
		httpClient: createHttpClient(cfg),
		semaphore:  newSemaphore(cfg.MaxConcurrentRequests()),
		stat:       new(statistics),
	}
}
