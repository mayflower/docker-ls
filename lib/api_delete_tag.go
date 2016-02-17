package lib

import (
	"fmt"
	"net/http"
)

func (r *registryApi) DeleteTag(repository, reference string) (err error) {
	response, err := r.connector.Delete(r.endpointUrl(fmt.Sprintf("/v2/%s/manifests/%s", repository, reference)), "")

	if err != nil {
		return
	}

	switch response.StatusCode {
	case http.StatusForbidden, http.StatusUnauthorized:
		err = genericAuthorizationError

	case http.StatusNotFound:
		err = newNotFoundError(fmt.Sprintf("%s:%s : no such repository or reference", repository, reference))

	case http.StatusBadRequest:
		err = newInvalidRequestError("invalid request --- make sure that your reference is a content digest")

	case http.StatusAccepted:

	default:
		err = newInvalidStatusCodeError(response.StatusCode)
	}

	return
}
