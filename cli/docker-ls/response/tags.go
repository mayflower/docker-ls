package response

import (
	"sync"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type tagL0 string

type tagsL0 struct {
	RepositoryName string  `yaml:"repository"`
	Tags           []tagL0 `yaml:"tags"`
	mutex          sync.Mutex
}

func (t *tagsL0) AddTag(tag lib.Tag) {
	t.mutex.Lock()
	t.Tags = append(t.Tags, tagL0(tag.Name()))
	t.mutex.Unlock()
}

func NewTagsL0(repositoryName string) *tagsL0 {
	return &tagsL0{
		RepositoryName: repositoryName,
	}
}
