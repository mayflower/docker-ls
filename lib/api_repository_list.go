package lib

import (
	"net/http"
)

type repositoryListResponse struct {
	repositories chan Repository
	err          error
}

func (r *repositoryListResponse) Repositories() <-chan Repository {
	return (r.repositories)
}

func (r *repositoryListResponse) LastError() error {
	return r.err
}

func (r *repositoryListResponse) setLastError(err error) {
	r.err = err
}

func (r *repositoryListResponse) close() {
	close(r.repositories)
}

type repositoryListJsonResponse struct {
	Repositories []string `json:"repositories"`
}

func (r *repositoryListJsonResponse) validate() error {
	if r.Repositories == nil {
		return genericMalformedResponseError
	}

	return nil
}

type repositoryListRequestContext struct{}

func (r *repositoryListRequestContext) path() string {
	return "v2/_catalog"
}

func (r *repositoryListRequestContext) validateApiResponse(response *http.Response, initialRequest bool) error {
	switch response.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden:
		return genericAuthorizationError

	case http.StatusNotFound:
		if initialRequest {
			return NotImplementedByRemoteError("registry does not implement repository listings")
		} else {
			return newInvalidStatusCodeError(response.StatusCode)
		}

	case http.StatusOK:
		return nil

	default:
		return newInvalidStatusCodeError(response.StatusCode)
	}
}

func (r *repositoryListRequestContext) processPartialResponse(response paginatedRequestResponse, apiResponse interface{}) {
	for _, repositoryName := range apiResponse.(*repositoryListJsonResponse).Repositories {
		response.(*repositoryListResponse).repositories <- newRepository(repositoryName)
	}
}

func (r *repositoryListRequestContext) createResponse(api *registryApi) paginatedRequestResponse {
	return &repositoryListResponse{
		repositories: make(chan Repository, api.pageSize()),
	}
}

func (r *repositoryListRequestContext) createJsonResponse() validatable {
	return new(repositoryListJsonResponse)
}

func (r *repositoryListRequestContext) tokenCacheHint() string {
	return cacheHintRegistryList()
}

func (r *repositoryListRequestContext) getHeaders() map[string]string {
	return nil
}

func (r *registryApi) ListRepositories() RepositoryListResponse {
	return r.paginatedRequest(new(repositoryListRequestContext)).(*repositoryListResponse)
}
