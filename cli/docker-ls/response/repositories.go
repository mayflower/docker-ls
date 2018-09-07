package response

import (
	"sort"
	"sync"

	"github.com/mayflower/docker-ls/lib"
)

type RepositoryL0 string

type RepositoryCollectionL0 []RepositoryL0

func (r RepositoryCollectionL0) Len() int {
	return len(r)
}

func (r RepositoryCollectionL0) Less(i, j int) bool {
	return string(r[i]) < string(r[j])
}

func (r RepositoryCollectionL0) Swap(i, j int) {
	tmp := r[i]
	r[i] = r[j]
	r[j] = tmp
}

type RepositoriesL0 struct {
	Repositories []RepositoryL0 `yaml:"repositories",json:"repositories"`
	mutex        sync.Mutex
}

func (r *RepositoriesL0) AddRepository(repo lib.Repository) {
	r.mutex.Lock()
	r.Repositories = append(r.Repositories, RepositoryL0(repo.Name()))
	r.mutex.Unlock()
}

func (r *RepositoriesL0) Sort() {
	r.mutex.Lock()
	sort.Sort(RepositoryCollectionL0(r.Repositories))
	r.mutex.Unlock()
}

func NewRepositoriesL0() *RepositoriesL0 {
	return new(RepositoriesL0)
}

type RepositoryL1 *TagsL0

type RepositoryCollectionL1 []RepositoryL1

func (r RepositoryCollectionL1) Len() int {
	return len(r)
}

func (r RepositoryCollectionL1) Less(i, j int) bool {
	return string(r[i].Repository) < string(r[j].Repository)
}

func (r RepositoryCollectionL1) Swap(i, j int) {
	tmp := r[i]
	r[i] = r[j]
	r[j] = tmp
}

type RepositoriesL1 struct {
	Repositories []RepositoryL1 `yaml:"repositories",json:"repositories"`
	mutex        sync.Mutex
}

func (r *RepositoriesL1) AddTags(tags *TagsL0) {
	r.mutex.Lock()
	r.Repositories = append(r.Repositories, RepositoryL1(tags))
	r.mutex.Unlock()
}

func (r *RepositoriesL1) Sort() {
	r.mutex.Lock()
	sort.Sort(RepositoryCollectionL1(r.Repositories))
	r.mutex.Unlock()

	for _, repository := range r.Repositories {
		(*TagsL0)(repository).Sort()
	}
}

func NewRepositoriesL1() *RepositoriesL1 {
	return new(RepositoriesL1)
}
