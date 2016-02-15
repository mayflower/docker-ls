package lib

import (
	"fmt"
)

type AutorizationError string
type NotImplementedByRemoteError string
type MalformedResponseError string
type InvalidStatusCodeError string

var genericAuthorizationError AutorizationError = "autorization failed"
var genericMalformedResponseError MalformedResponseError = "malformed response"

func (e AutorizationError) Error() string {
	return string(e)
}

func (e NotImplementedByRemoteError) Error() string {
	return string(e)
}

func (e MalformedResponseError) Error() string {
	return string(e)
}

func (e InvalidStatusCodeError) Error() string {
	return string(e)
}

func newInvalidStatusCodeError(code int) error {
	return InvalidStatusCodeError(fmt.Sprintf("invalid API response status %d", code))
}
