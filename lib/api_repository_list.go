package lib

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type RepositoryListResponse struct {
	repositories chan *Repository
	err          error
}

type repositoryListJsonResponse struct {
	Repositories *[]string `json:"repositories"`
}

func (r *RepositoryListResponse) Repositories() <-chan *Repository {
	return (r.repositories)
}

func (r *RepositoryListResponse) LastError() error {
	return r.err
}

func (r *RegistryApi) executeListRequest(url *url.URL) (response *http.Response, close bool, err error) {
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
		err = NotImplementedByRemoteError("registry does not implement repository listings")
		return

	case http.StatusOK:

	default:
		err = newInvalidStatusCodeError(response.StatusCode)
		return
	}

	return
}

func (r *RegistryApi) iterateRepositoryList(lastApiResponse *http.Response, listResponse *RepositoryListResponse) (apiResponse *http.Response, more bool, err error) {
	requestUrl := r.endpointUrl("v2/_catalog")

	if lastApiResponse != nil {
		linkHeader := lastApiResponse.Header.Get("link")

		if linkHeader != "" {
			// This is a hack to work around what looks like a bug in the registry:
			// the supplied link URL currently lacks scheme and host
			scheme, host := requestUrl.Scheme, requestUrl.Host

			requestUrl, err = parseLinkToNextHeader(linkHeader)
			requestUrl.Scheme = scheme
			requestUrl.Host = host
		}

		if err != nil {
			return
		}
	} else {
		queryParams := requestUrl.Query()
		queryParams.Set("n", strconv.Itoa(r.pageSize))
		requestUrl.RawQuery = queryParams.Encode()
	}

	apiResponse, needsClose, err := r.executeListRequest(requestUrl)

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

func (r *RegistryApi) ListRepositories() (response *RepositoryListResponse, err error) {
	response = &RepositoryListResponse{
		repositories: make(chan *Repository, r.pageSize),
	}

	var apiResponse *http.Response
	more := true

	go func() {
		for more {
			apiResponse, more, err = r.iterateRepositoryList(apiResponse, response)

			if err != nil {
				response.err = err
				break
			}
		}

		close(response.repositories)
	}()

	return
}
