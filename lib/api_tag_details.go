package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type layerDetails string

func (l layerDetails) ContentDigest() string {
	return string(l)
}

type tagDetails struct {
	manifest      *parsedManifest
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
	return t.manifest.Name
}

func (t *tagDetails) TagName() string {
	return t.manifest.Tag
}

func (t *tagDetails) Layers() []LayerDetails {
	return t.layers
}

func (t *tagDetails) setLayers(layers []parsedLayer) {
	t.layers = make([]LayerDetails, 0, len(layers))

	for _, layer := range layers {
		t.layers = append(t.layers, layerDetails(layer.BlobSum))
	}
}

func (r *registryApi) GetTagDetails(repository, reference string) (details TagDetails, err error) {
	url := r.endpointUrl(fmt.Sprintf("v2/%s/manifests/%s", repository, reference))

	apiResponse, err := r.connector.Get(url, cacheHintTagDetails(repository))

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
		err = newNotFoundError(fmt.Sprintf("%s:%s : no such repository or reference", repository, reference))

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

	var manifest parsedManifest
	err = json.Unmarshal(bodyBuffer.Bytes(), &manifest)
	if err != nil {
		return
	}

	_details := &tagDetails{
		manifest:      &manifest,
		rawManifest:   rawManifest,
		contentDigest: apiResponse.Header.Get("docker-content-digest"),
	}
	_details.setLayers(manifest.Layers)
	details = _details

	return
}
