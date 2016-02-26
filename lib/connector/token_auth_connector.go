package connector

import (
	"net/http"
	"net/url"
	"strings"

	"git.mayflower.de/vaillant-team/docker-ls/lib/auth"
)

type tokenAuthConnector struct {
	cfg           TokenAuthConfig
	httpClient    *http.Client
	authenticator auth.Authenticator
	semaphore     chan int
	tokenCache    *tokenCache
	stat          *statistics
	basicAuth     bool
}

func (r *tokenAuthConnector) AquireLock() {
	r.semaphore <- 1
}

func (r *tokenAuthConnector) ReleaseLock() {
	_ = <-r.semaphore
}

func (r *tokenAuthConnector) Delete(url *url.URL, hint string) (*http.Response, error) {
	return r.Request("DELETE", url, hint)
}

func (r *tokenAuthConnector) Get(url *url.URL, hint string) (*http.Response, error) {
	return r.Request("GET", url, hint)
}

func (r *tokenAuthConnector) Request(method string, url *url.URL, hint string) (response *http.Response, err error) {
	r.AquireLock()
	defer r.ReleaseLock()

	r.stat.Request()

	var token auth.Token
	request, err := http.NewRequest(method, url.String(), strings.NewReader(""))

	if err != nil {
		return
	}

	if hint != "" {
		if token = r.tokenCache.Get(hint); token != nil {
			r.stat.CacheHitAtApiLevel()
		} else {
			r.stat.CacheMissAtApiLevel()
		}
	}

	resp, err := r.attemptRequestWithToken(request, token)

	if err != nil || resp.StatusCode != http.StatusUnauthorized {
		response = resp
		return
	}

	if token != nil {
		r.stat.CacheFailAtApiLevel()
	}

	if resp.Close {
		resp.Body.Close()
	}

	challenge, err := auth.ParseChallenge(resp.Header.Get("www-authenticate"))

	if err != nil {
		return
	}

	token, err = r.authenticator.Authenticate(challenge, false)

	if err != nil {
		return
	}

	if token != nil {
		if token.Fresh() {
			r.stat.CacheMissAtAuthLevel()
		} else {
			r.stat.CacheHitAtAuthLevel()
		}
	}

	response, err = r.attemptRequestWithToken(request, token)

	if err == nil &&
		response.StatusCode == http.StatusUnauthorized &&
		!token.Fresh() {

		r.stat.CacheFailAtAuthLevel()

		token, err = r.authenticator.Authenticate(challenge, true)

		if err == nil {
			return
		}

		response, err = r.attemptRequestWithToken(request, token)
	}

	if hint != "" && err == nil && response.StatusCode != http.StatusUnauthorized {
		r.tokenCache.Set(hint, token)
	}

	return
}

func (r *tokenAuthConnector) attemptRequestWithToken(request *http.Request, token auth.Token) (*http.Response, error) {
	if token != nil {
		request.Header.Set("Authorization", "Bearer "+token.Value())
	}

	return r.httpClient.Do(request)
}

func (r *tokenAuthConnector) GetStatistics() Statistics {
	return r.stat
}

func NewTokenAuthConnector(cfg TokenAuthConfig) *tokenAuthConnector {
	connector := tokenAuthConnector{
		cfg:        cfg,
		httpClient: http.DefaultClient,
		semaphore:  make(chan int, cfg.MaxConcurrentRequests()),
		tokenCache: newTokenCache(),
		stat:       new(statistics),
	}

	connector.authenticator = auth.NewAuthenticator(connector.httpClient, cfg.Credentials())

	return &connector
}
