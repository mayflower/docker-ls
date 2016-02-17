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
	RepositoryName string  `yaml:"repository",json:"repository"`
	Tags           []TagL0 `yaml:"tags",json:"tags"`
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

type TagL1 *TagDetailsL0

type TagCollectionL1 []TagL1

func (t TagCollectionL1) Len() int {
	return len(t)
}

func (t TagCollectionL1) Less(i, j int) bool {
	return string(t[i].TagName) < string(t[j].TagName)
}

func (t TagCollectionL1) Swap(i, j int) {
	tmp := t[i]
	t[i] = t[j]
	t[j] = tmp
}

type TagsL1 struct {
	RepositoryName string `yaml:"repository",json:"repository"`
	Tags           []TagL1
	mutex          sync.Mutex
}

func (t *TagsL1) AddTag(tag lib.TagDetails) {
	t.mutex.Lock()
	t.Tags = append(t.Tags, NewTagDetailsL0(tag, false))
	t.mutex.Unlock()
}

func (t *TagsL1) Sort() {
	t.mutex.Lock()
	sort.Sort(TagCollectionL1(t.Tags))
	t.mutex.Unlock()
}

func NewTagsL1(repositoryName string) *TagsL1 {
	return &TagsL1{
		RepositoryName: repositoryName,
	}
}
