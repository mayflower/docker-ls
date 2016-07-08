package lib

import (
	"net/url"

	"github.com/mayflower/docker-ls/lib/connector"
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
	GetRegistryUrl() *url.URL
	ListRepositories() RepositoryListResponse
	ListTags(repositoryName string) TagListResponse
	GetTagDetails(ref Refspec, manifestVersion uint) (TagDetails, error)
	DeleteTag(ref Refspec) error
	GetStatistics() connector.Statistics
}
