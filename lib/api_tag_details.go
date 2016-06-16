package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type layerDetails string

func (l layerDetails) ContentDigest() string {
	return string(l)
}

type tagDetails struct {
	name          string
	tag           string
	rawManifest   interface{}
	contentDigest string
	layers        []LayerDetails
}

func (t *tagDetails) RawManifest() interface{} {
	return t.rawManifest
}

func (t *tagDetails) ContentDigest() string {
	return t.contentDigest
}

func (t *tagDetails) RepositoryName() string {
	return t.name
}

func (t *tagDetails) TagName() string {
	return t.tag
}

func (t *tagDetails) Layers() []LayerDetails {
	return t.layers
}

func (t *tagDetails) setLayers(layers []parsedLayer) {
	t.layers = make([]LayerDetails, 0, len(layers))

	for _, layer := range layers {
		t.layers = append(t.layers, layerDetails(layer.blobSum()))
	}
}

func (r *registryApi) GetTagDetails(ref Refspec, manifestVersion uint) (details TagDetails, err error) {
	url := r.endpointUrl(fmt.Sprintf("v2/%s/manifests/%s", ref.Repository(), ref.Reference()))

	headers, err := r.getHeadersForManifestVersion(manifestVersion)
	if err != nil {
		return
	}

	apiResponse, err := r.connector.Get(
		url,
		headers,
		cacheHintTagDetails(ref.Repository()),
	)

	if err != nil {
		return
	}

	if apiResponse.Close {
		defer apiResponse.Body.Close()
	}

	switch apiResponse.StatusCode {
	case http.StatusForbidden, http.StatusUnauthorized:
		err = genericAuthorizationError

	case http.StatusNotFound:
		err = newNotFoundError(fmt.Sprintf("%v : no such repository or reference", ref))

	case http.StatusOK:

	default:
		err = newInvalidStatusCodeError(apiResponse.StatusCode)
	}

	if err != nil {
		return
	}

	bodyBuffer := bytes.Buffer{}
	_, err = io.Copy(&bodyBuffer, apiResponse.Body)
	if err != nil {
		return
	}

	var rawManifest interface{}
	err = json.Unmarshal(bodyBuffer.Bytes(), &rawManifest)
	if err != nil {
		return
	}

	parsedManifest, err := parseManifest(bodyBuffer.Bytes())
	if err != nil {
		return
	}

	_details := &tagDetails{
		rawManifest:   rawManifest,
		name:          ref.Repository(),
		tag:           ref.Reference(),
		contentDigest: apiResponse.Header.Get("docker-content-digest"),
	}

	_details.setLayers(parsedManifest.layers())
	details = _details

	return
}

func (r *registryApi) getHeadersForManifestVersion(version uint) (headers map[string]string, err error) {
	switch version {
	case 1:
		headers = nil

	case 2:
		headers = map[string]string{
			"accept": "application/vnd.docker.distribution.manifest.v2+json",
		}

	default:
		err = errors.New("invalid manifest version")
	}

	return
}
