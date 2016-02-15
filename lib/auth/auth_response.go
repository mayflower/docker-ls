package auth

import (
	"encoding/json"
	"errors"
	"io"
)

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
