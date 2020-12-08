package auth

type RegistryCredentials interface {
	User() string
	Password() string
	IdentityToken() string
	SetPassword(password string)
	SetUser(user string)
}
