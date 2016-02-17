package auth

import (
	"errors"
	"fmt"
	"net/http"
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
		err = errors.New(fmt.Sprintf("authentication failed with code %d", authResponse.StatusCode))
		return
	}

	decodedResponse, err := decodeAuthResponse(authResponse.Body)

	if err != nil {
		return
	}

	a.cache.Set(c, decodedResponse)
	t = newToken(decodedResponse.Token, true)

	return
}

func NewAuthenticator(client *http.Client, credentials RegistryCredentials) Authenticator {
	return &authenticator{
		httpClient:  client,
		credentials: credentials,
		cache:       newTokenCache(),
	}
}
