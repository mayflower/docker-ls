package lib

import (
	"encoding/json"
	"io"
)

type parsedLayer struct {
	BlobSum string `json:"blobSum"`
}

type parsedManifest struct {
	Name   string        `json:"name"`
	Tag    string        `json:"tag"`
	Layers []parsedLayer `json:"fsLayers"`
}

func parseManifest(data io.Reader) (manifest *parsedManifest, err error) {
	decoder := json.NewDecoder(data)
	err = decoder.Decode(&manifest)

	return
}
