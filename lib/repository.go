package lib

type Repository struct {
	name string
}

func (r *Repository) Name() string {
	return r.name
}

func newRepository(name string) *Repository {
	return &Repository{
		name: name,
	}
}
