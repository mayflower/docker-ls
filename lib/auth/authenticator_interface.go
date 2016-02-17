package auth

type Authenticator interface {
	Authenticate(challenge *Challenge, ignoreCached bool) (Token, error)
}
