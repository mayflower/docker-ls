package response

import (
	"sync"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type TagL0 string

type TagsL0 struct {
	RepositoryName string  `yaml:"repository"`
	Tags           []TagL0 `yaml:"tags"`
	mutex          sync.Mutex
}

func (t *TagsL0) AddTag(tag lib.Tag) {
	t.mutex.Lock()
	t.Tags = append(t.Tags, TagL0(tag.Name()))
	t.mutex.Unlock()
}

func NewTagsL0(repositoryName string) *TagsL0 {
	return &TagsL0{
		RepositoryName: repositoryName,
	}
}
