package lib

type repository struct {
	name string
}

func (r *repository) Name() string {
	return r.name
}

func newRepository(name string) *repository {
	return &repository{
		name: name,
	}
}
