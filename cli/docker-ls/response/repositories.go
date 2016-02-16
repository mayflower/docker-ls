package response

import (
	"sync"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type RepositoryL0 string

type RepositoriesL0 struct {
	Repositories []RepositoryL0 `yaml:"repositories"`
	mutex        sync.Mutex
}

func (r *RepositoriesL0) AddRepository(repo lib.Repository) {
	r.mutex.Lock()
	r.Repositories = append(r.Repositories, RepositoryL0(repo.Name()))
	r.mutex.Unlock()
}

func NewRepositoriesL0() *RepositoriesL0 {
	return new(RepositoriesL0)
}

type RepositoryL1 *TagsL0

type RepositoriesL1 struct {
	Tags  []RepositoryL1 `yaml:"repositories"`
	mutex sync.Mutex
}

func (r *RepositoriesL1) AddTags(tags *TagsL0) {
	r.mutex.Lock()
	r.Tags = append(r.Tags, RepositoryL1(tags))
	r.mutex.Unlock()
}

func NewRepositoriesL1() *RepositoriesL1 {
	return new(RepositoriesL1)
}
