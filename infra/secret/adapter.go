package secret

type SecretAdapter interface {
	GetSecret(key string) string
	InitEnvs()
}
