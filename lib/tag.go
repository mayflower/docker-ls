package lib

type tag struct {
	name           string
	repositoryName string
}

func (t *tag) Name() string {
	return t.name
}

func (t *tag) RepositoryName() string {
	return t.repositoryName
}

func newTag(name, repositoryName string) *tag {
	return &tag{
		name:           name,
		repositoryName: repositoryName,
	}
}
