package response

import (
	"sync"

	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type repositoryL0 string

type repositoriesL0 struct {
	Repositories []repositoryL0 `yaml:"repositories"`
	mutex        sync.Mutex
}

func (r *repositoriesL0) AddRepository(repo lib.Repository) {
	r.mutex.Lock()
	r.Repositories = append(r.Repositories, repositoryL0(repo.Name()))
	r.mutex.Unlock()
}

func NewRepositoriesL0() *repositoriesL0 {
	return new(repositoriesL0)
}
