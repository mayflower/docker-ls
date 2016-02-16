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

type TagDetails interface {
	RawManifest() interface{}
	ContentDigest() string
	RepositoryName() string
	TagName() string
}

type RegistryApi interface {
	ListRepositories() RepositoryListResponse
	ListTags(repositoryName string) TagListResponse
	GetTagDetails(repository, reference string) (TagDetails, error)
}
