package lib

import (
	"encoding/json"
	"io"
)

type parsedManifest struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func parseManifest(data io.Reader) (manifest *parsedManifest, err error) {
	decoder := json.NewDecoder(data)
	err = decoder.Decode(&manifest)

	return
}
