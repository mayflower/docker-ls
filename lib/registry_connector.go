package lib

import (
	"net/http"
	"strings"

	"git.mayflower.de/vaillant-team/docker-ls/lib/auth"
)

type RegistryConnector struct {
	cfg           Config
	httpClient    *http.Client
	authenticator auth.AuthenticatorInterface
}

func (r *RegistryConnector) Get(path string) (response *http.Response, err error) {
	endpoint := strings.TrimRight(r.cfg.registryUrl.String(), "/") + "/" + strings.TrimLeft(path, "/")

	request, err := http.NewRequest("GET", endpoint, strings.NewReader(""))

	if err != nil {
		return
	}

	resp, err := r.httpClient.Do(request)

	if err != nil || resp.StatusCode != http.StatusUnauthorized {
		response = resp
		return
	}

	if resp.Close {
		resp.Body.Close()
	}

	challenge, err := auth.ParseChallenge(resp.Header.Get("www-authenticate"))

	if err != nil {
		return
	}

	token, err := r.authenticator.PerformRequest(challenge)

	if err != nil {
		return
	}

	request.Header.Set("Authorization", "Bearer "+token)

	response, err = r.httpClient.Do(request)

	return
}

func NewRegistryConnector(cfg Config) *RegistryConnector {
	connector := RegistryConnector{
		cfg:        cfg,
		httpClient: http.DefaultClient,
	}

	connector.authenticator = auth.NewAuthenticator(connector.httpClient, &cfg.credentials)

	return &connector
}
