package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type authenticator struct {
	httpClient  *http.Client
	credentials RegistryCredentials
	cache       *tokenCache
}

func (a *authenticator) Authenticate(c *Challenge, ignoreCached bool) (t Token, err error) {
	if !ignoreCached {
		value := a.cache.Get(c)

		if value != "" {
			t = newToken(value, false)
			return
		}
	}

	// Try OAuth2 authentication first, and then legacy JWT tokens.
	var decodedResponse authResponse
	if refreshToken := a.credentials.IdentityToken(); refreshToken != "" {
		decodedResponse, err = a.fetchTokenOAuth2(c, refreshToken)
	} else {
		decodedResponse, err = a.fetchTokenJWT(c)
	}
	if err != nil {
		return
	}

	a.cache.Set(c, decodedResponse)
	t = newToken(decodedResponse.Token, true)

	return
}

func (a *authenticator) fetchTokenJWT(c *Challenge) (decodedResponse authResponse, err error) {
	requestUrl := c.buildRequestUrl()
	authRequest, err := http.NewRequest("GET", requestUrl.String(), strings.NewReader(""))
	if err != nil {
		return
	}

	username := a.credentials.User()
	password := a.credentials.Password()
	if username != "" || password != "" {
		authRequest.SetBasicAuth(username, password)
	}

	authResponse, err := a.httpClient.Do(authRequest)
	if err != nil {
		return
	}

	if authResponse.Close {
		defer authResponse.Body.Close()
	}

	if authResponse.StatusCode != http.StatusOK {
		err = fmt.Errorf("authentication against auth server failed with code %d", authResponse.StatusCode)
		return
	}

	decodedResponse, err = decodeAuthResponse(authResponse.Body)
	return
}

func (a *authenticator) fetchTokenOAuth2(c *Challenge, refreshToken string) (decodedResponse authResponse, err error) {
	authResponse, err := a.httpClient.PostForm(c.realm.String(), url.Values{
		"grant_type":    {"refresh_token"},
		"service":       {c.service},
		"client_id":     {"docker-ls"},
		"scope":         {strings.Join(c.scope, " ")},
		"refresh_token": {refreshToken},
	})
	if err != nil {
		return
	}

	if authResponse.StatusCode != http.StatusOK {
		err = fmt.Errorf("OAuth2 authentication against auth server failed with code %d", authResponse.StatusCode)
		return
	}

	decodedResponse, err = decodeAuth2Response(authResponse.Body)
	return
}

func NewAuthenticator(client *http.Client, credentials RegistryCredentials) Authenticator {
	return &authenticator{
		httpClient:  client,
		credentials: credentials,
		cache:       newTokenCache(),
	}
}
