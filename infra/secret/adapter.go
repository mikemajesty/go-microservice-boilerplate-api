package secret

type Adapter interface {
	GetSecret(key string) string
	InitEnvs()
}
