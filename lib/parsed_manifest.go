package lib

import (
	"encoding/json"
	"errors"
)

type parsedLayer interface {
	blobSum() string
}

type parsedManifest interface {
	layers() []parsedLayer
}

type parsedLayerV1 struct {
	BlobSum string `json:"blobSum"`
}

type parsedManifestV1 struct {
	Layers []parsedLayerV1 `json:"fsLayers"`
}

type parsedLayerV2 struct {
	Digest string `json:"digest"`
}

type parsedManifestV2 struct {
	Layers []parsedLayerV2 `json:"layers"`
}

func (p *parsedLayerV1) blobSum() string {
	return p.BlobSum
}

func (m *parsedManifestV1) layers() (layers []parsedLayer) {
	layers = make([]parsedLayer, 0, len(m.Layers))

	for i := 0; i < len(m.Layers); i++ {
		layers = append(layers, &m.Layers[i])
	}

	return
}

func (p *parsedLayerV2) blobSum() string {
	return p.Digest
}

func (m *parsedManifestV2) layers() (layers []parsedLayer) {
	layers = make([]parsedLayer, 0, len(m.Layers))

	for i := 0; i < len(m.Layers); i++ {
		layers = append(layers, &m.Layers[i])
	}

	return
}

func parseManifest(data []byte) (manifest parsedManifest, err error) {
	var id struct {
		SchemaVersion int `json:"schemaVersion"`
	}

	err = json.Unmarshal(data, &id)
	if err != nil {
		return
	}

	switch id.SchemaVersion {
	case 1:
		var parsedData parsedManifestV1
		err = json.Unmarshal(data, &parsedData)
		manifest = &parsedData

	case 2:
		var parsedData parsedManifestV2
		err = json.Unmarshal(data, &parsedData)
		manifest = &parsedData

	default:
		err = errors.New("unknown manifest schema version")
	}

	return
}
