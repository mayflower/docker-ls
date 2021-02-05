package lib

import (
	"fmt"
	"log"
	"net/http"
)

type tagListResponse struct {
	tags chan Tag
	err  error
}

func (t *tagListResponse) Tags() <-chan Tag {
	return t.tags
}

func (t *tagListResponse) LastError() error {
	return t.err
}

func (r *tagListResponse) setLastError(err error) {
	r.err = err
}

func (r *tagListResponse) close() {
	close(r.tags)
}

type tagListJsonResponse struct {
	RepositoryName string   `json:"name"`
	Tags           []string `json:"tags"`
}

func (r *tagListJsonResponse) validate() error {
	if r.RepositoryName == "" {
		return genericMalformedResponseError
	}

	return nil
}

type tagListRequestContext struct {
	repositoryName string
}

func (r *tagListRequestContext) path() string {
	return fmt.Sprintf("v2/%s/tags/list", r.repositoryName)
}

func (r *tagListRequestContext) validateApiResponse(response *http.Response, initialRequest bool) error {
	log.Printf("api response: %#v", response)

	switch response.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden:
		return genericAuthorizationError

	case http.StatusNotFound:
		if initialRequest {
			return newNotFoundError(fmt.Sprintf("%s: no such repository", r.repositoryName))
		} else {
			return newInvalidStatusCodeError(response.StatusCode)
		}

	case http.StatusOK:
		return nil

	default:
		return newInvalidStatusCodeError(response.StatusCode)
	}
}

func (r *tagListRequestContext) processPartialResponse(response paginatedRequestResponse, apiResponse interface{}) {
	for _, tagName := range apiResponse.(*tagListJsonResponse).Tags {
		response.(*tagListResponse).tags <- newTag(tagName, r.repositoryName)
	}
}

func (r *tagListRequestContext) createResponse(api *registryApi) paginatedRequestResponse {
	return &tagListResponse{
		tags: make(chan Tag, api.pageSize()),
	}
}

func (r *tagListRequestContext) createJsonResponse() validatable {
	return new(tagListJsonResponse)
}

func (r *tagListRequestContext) tokenCacheHint() string {
	return cacheHintTagList(r.repositoryName)
}

func (r *tagListRequestContext) getHeaders() map[string]string {
	return nil
}

func (r *registryApi) ListTags(repositoryName string) TagListResponse {
	ctx := tagListRequestContext{
		repositoryName: repositoryName,
	}

	return r.paginatedRequest(&ctx).(*tagListResponse)
}
