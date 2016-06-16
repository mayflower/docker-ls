package lib

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type validatable interface {
	validate() error
}

type paginatedRequestResponse interface {
	setLastError(error)
	close()
}

type paginatedRequestContext interface {
	path() string
	tokenCacheHint() string
	validateApiResponse(response *http.Response, initialRequest bool) error
	processPartialResponse(response paginatedRequestResponse, apiResponse interface{})
	createResponse(api *registryApi) paginatedRequestResponse
	createJsonResponse() validatable
	getHeaders() map[string]string
}

func (r *registryApi) executePaginatedRequest(ctx paginatedRequestContext, url *url.URL, initialRequest bool) (response *http.Response, close bool, err error) {
	response, err = r.connector.Get(url, ctx.getHeaders(), ctx.tokenCacheHint())

	if err != nil {
		return
	}

	if err == nil {
		close = response.Close
		err = ctx.validateApiResponse(response, initialRequest)
	}

	return
}

func (r *registryApi) iteratePaginatedRequest(
	ctx paginatedRequestContext,
	lastApiResponse *http.Response,
	response paginatedRequestResponse,
) (
	apiResponse *http.Response,
	more bool,
	err error,
) {
	requestUrl, err := r.paginatedRequestEndpointUrl(ctx.path(), lastApiResponse)

	if err != nil {
		return
	}

	apiResponse, needsClose, err := r.executePaginatedRequest(ctx, requestUrl, lastApiResponse == nil)

	if needsClose {
		defer apiResponse.Body.Close()
	}

	if err != nil {
		return
	}

	more = apiResponse.Header.Get("link") != ""

	jsonResponse := ctx.createJsonResponse()
	decoder := json.NewDecoder(apiResponse.Body)
	err = decoder.Decode(&jsonResponse)

	if err != nil {
		return
	}

	err = jsonResponse.validate()

	if err != nil {
		return
	}

	ctx.processPartialResponse(response, jsonResponse)

	return
}

func (r *registryApi) paginatedRequest(ctx paginatedRequestContext) (response paginatedRequestResponse) {
	response = ctx.createResponse(r)

	go func() {
		var apiResponse *http.Response
		var err error
		more := true

		for more {
			apiResponse, more, err = r.iteratePaginatedRequest(ctx, apiResponse, response)

			if err != nil {
				response.setLastError(err)
				break
			}
		}

		response.close()
	}()

	return
}
