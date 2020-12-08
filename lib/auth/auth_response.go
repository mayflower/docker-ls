package auth

import (
	"encoding/json"
	"errors"
	"io"
)

// authResponse stores an JWT token or OAuth2 Access token.
type authResponse struct {
	Token string `json:"token"`
}

func decodeAuthResponse(serverResponse io.Reader) (response authResponse, err error) {
	decoder := json.NewDecoder(serverResponse)

	err = decoder.Decode(&response)

	if err == nil && response.Token == "" {
		err = errors.New("malformed auth server response")
	}

	return
}

// auth2Response stores an OAuth2 response.
// https://docs.docker.com/registry/spec/auth/oauth/
type auth2Response struct {
	AccessToken string `json:"access_token"`
}

func decodeAuth2Response(serverResponse io.Reader) (response authResponse, err error) {
	decoder := json.NewDecoder(serverResponse)

	var oauth2Response auth2Response
	err = decoder.Decode(&oauth2Response)
	response.Token = oauth2Response.AccessToken

	if err == nil && response.Token == "" {
		err = errors.New("malformed auth server response")
	}

	return
}
