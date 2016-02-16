package lib

type Repository interface {
	Name() string
}

type RepositoryListResponse interface {
	Repositories() <-chan Repository
	LastError() error
}

type Tag interface {
	Name() string
	RepositoryName() string
}

type TagListResponse interface {
	Tags() <-chan Tag
	LastError() error
}

type RegistryApi interface {
	ListRepositories() RepositoryListResponse
	ListTags(repositoryName string) TagListResponse
}
