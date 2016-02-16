package auth

type Authenticator interface {
	PerformRequest(*Challenge) (token string, err error)
}
