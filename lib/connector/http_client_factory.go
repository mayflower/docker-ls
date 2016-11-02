package connector

import (
	"crypto/tls"
	"net/http"
)

func createHttpClient(cfg Config) *http.Client {
	var tlsConfig *tls.Config
	if cfg.AllowInsecure() {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return &http.Client{
		Transport: &http.Transport{
                Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: tlsConfig,
		},
	}
}
