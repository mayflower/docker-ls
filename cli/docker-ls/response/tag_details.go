package response

import (
	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type tagDetailsL0 struct {
	RepositoryName string `yaml:"repository"`
	TagName        string `yaml:"tagName"`
	ContentDigest  string `yaml:"digest"`
}

func NewTagDetailsL0(tag lib.TagDetails) *tagDetailsL0 {
	return &tagDetailsL0{
		RepositoryName: tag.RepositoryName(),
		TagName:        tag.TagName(),
		ContentDigest:  tag.ContentDigest(),
	}
}
