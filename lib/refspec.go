package lib

import (
	"errors"
	"flag"
	"strings"
)

type Refspec interface {
	flag.Value

	Repository() string
	Reference() string
}

type refspec struct {
	repository string
	reference  string
}

func (r *refspec) Repository() string {
	return r.repository
}

func (r *refspec) Reference() string {
	return r.reference
}

func (r *refspec) String() string {
	return r.repository + ":" + r.reference
}

func (r *refspec) Set(value string) (err error) {
	if strings.Contains(value, ":") {
		pieces := strings.SplitN(value, ":", 2)
		r.repository, r.reference = pieces[0], pieces[1]
	} else {
		err = errors.New("invalid refspec")
	}

	return
}

func EmptyRefspec() Refspec {
	return new(refspec)
}

func NewRefspec(repository, reference string) Refspec {
	return &refspec{
		repository: repository,
		reference:  reference,
	}
}
