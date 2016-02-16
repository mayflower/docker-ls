package response

import (
	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type TagDetailsL0 struct {
	RepositoryName string `yaml:"repository"`
	TagName        string `yaml:"tagName"`
	ContentDigest  string `yaml:"digest"`
}

func NewTagDetailsL0(tag lib.TagDetails) *TagDetailsL0 {
	return &TagDetailsL0{
		RepositoryName: tag.RepositoryName(),
		TagName:        tag.TagName(),
		ContentDigest:  tag.ContentDigest(),
	}
}
