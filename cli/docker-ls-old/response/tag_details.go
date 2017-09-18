package response

import (
	"github.com/mayflower/docker-ls/lib"
)

type LayerL0 string

type TagDetailsL0 struct {
	RepositoryName string    `yaml:"repository,omitempty",json:"repository,omitempty"`
	TagName        string    `yaml:"tagName",json:"tagName"`
	ContentDigest  string    `yaml:"digest",json:"digest"`
	Layers         []LayerL0 `yaml:"layers",json:"layers"`
}

func NewTagDetailsL0(tag lib.TagDetails, includeRepository bool) *TagDetailsL0 {
	layers := make([]LayerL0, 0, len(tag.Layers()))
	for _, layer := range tag.Layers() {
		layers = append(layers, LayerL0(layer.ContentDigest()))
	}

	tagDetails := TagDetailsL0{
		TagName:       tag.TagName(),
		ContentDigest: tag.ContentDigest(),
		Layers:        layers,
	}

	if includeRepository {
		tagDetails.RepositoryName = tag.RepositoryName()
	}

	return &tagDetails
}
