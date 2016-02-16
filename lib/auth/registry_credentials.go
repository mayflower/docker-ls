package auth

type RegistryCredentials interface {
	User() string
	Password() string
}
