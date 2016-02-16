package lib

type Repository interface {
	Name() string
}

type RepositoryListResponse interface {
	Repositories() <-chan Repository
	LastError() error
}

type RegistryApi interface {
	ListRepositories() (RepositoryListResponse, error)
}
