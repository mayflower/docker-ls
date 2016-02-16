package response

import (
	"sort"
	"sync"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type TagL0 string

type TagCollectionL0 []TagL0

func (t TagCollectionL0) Len() int {
	return len(t)
}

func (t TagCollectionL0) Less(i, j int) bool {
	return string(t[i]) < string(t[j])
}

func (t TagCollectionL0) Swap(i, j int) {
	tmp := t[i]
	t[i] = t[j]
	t[j] = tmp
}

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

func (t *TagsL0) Sort() {
	t.mutex.Lock()
	sort.Sort(TagCollectionL0(t.Tags))
	t.mutex.Unlock()
}

func NewTagsL0(repositoryName string) *TagsL0 {
	return &TagsL0{
		RepositoryName: repositoryName,
	}
}
