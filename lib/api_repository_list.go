package lib

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type repositoryListResponse struct {
	repositories chan Repository
	err          error
}

type repositoryListJsonResponse struct {
	Repositories *[]string `json:"repositories"`
}

func (r *repositoryListResponse) Repositories() <-chan Repository {
	return (r.repositories)
}

func (r *repositoryListResponse) LastError() error {
	return r.err
}

func (r *registryApi) executeListRequest(url *url.URL, initialRequest bool) (response *http.Response, close bool, err error) {
	response, err = r.connector.Get(url)

	if err != nil {
		return
	}

	close = response.Close

	switch response.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden:
		err = genericAuthorizationError
		return

	case http.StatusNotFound:
		if initialRequest {
			err = NotImplementedByRemoteError("registry does not implement repository listings")
		} else {
			err = newInvalidStatusCodeError(response.StatusCode)
		}

		return

	case http.StatusOK:

	default:
		err = newInvalidStatusCodeError(response.StatusCode)
		return
	}

	return
}

func (r *registryApi) iterateRepositoryList(lastApiResponse *http.Response, listResponse *repositoryListResponse) (apiResponse *http.Response, more bool, err error) {
	requestUrl, err := r.paginatedRequestEndpointUrl("v2/_catalog", lastApiResponse)

	if err != nil {
		return
	}

	apiResponse, needsClose, err := r.executeListRequest(requestUrl, lastApiResponse == nil)

	if needsClose {
		defer apiResponse.Body.Close()
	}

	if err != nil {
		return
	}

	more = apiResponse.Header.Get("link") != ""

	var jsonResponse repositoryListJsonResponse
	decoder := json.NewDecoder(apiResponse.Body)
	err = decoder.Decode(&jsonResponse)

	if err != nil {
		return
	}

	if jsonResponse.Repositories == nil {
		err = genericMalformedResponseError
		return
	}

	for _, repositoryName := range *jsonResponse.Repositories {
		listResponse.repositories <- newRepository(repositoryName)
	}

	return
}

func (r *registryApi) ListRepositories() (response RepositoryListResponse, err error) {
	listResponse := &repositoryListResponse{
		repositories: make(chan Repository, r.pageSize()),
	}
	response = listResponse

	var apiResponse *http.Response
	apiResponse, more, err := r.iterateRepositoryList(apiResponse, listResponse)

	go func() {
		for more {
			var err error
			apiResponse, more, err = r.iterateRepositoryList(apiResponse, listResponse)

			if err != nil {
				listResponse.err = err
				break
			}
		}

		close(listResponse.repositories)
	}()

	return
}
