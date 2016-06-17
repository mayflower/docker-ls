package auth

type RegistryCredentials interface {
	User() string
	Password() string
	SetPassword(password string)
}
