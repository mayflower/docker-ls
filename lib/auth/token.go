package auth

type Token interface {
	Value() string
	Fresh() bool
}

type token struct {
	value string
	fresh bool
}

func (t *token) Value() string {
	return t.value
}

func (t *token) Fresh() bool {
	return t.fresh
}

func newToken(value string, fresh bool) Token {
	return &token{
		value: value,
		fresh: fresh,
	}
}
