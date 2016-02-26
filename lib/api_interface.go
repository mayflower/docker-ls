package lib

import (
	"git.mayflower.de/vaillant-team/docker-ls/lib/connector"
)

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

type LayerDetails interface {
	ContentDigest() string
}

type TagDetails interface {
	RawManifest() interface{}
	ContentDigest() string
	RepositoryName() string
	TagName() string
	Layers() []LayerDetails
}

type RegistryApi interface {
	ListRepositories() RepositoryListResponse
	ListTags(repositoryName string) TagListResponse
	GetTagDetails(ref Refspec) (TagDetails, error)
	DeleteTag(ref Refspec) error
	GetStatistics() connector.Statistics
}
