package auth

type RegistryCredentialsInterface interface {
	User() string
	Password() string
}
