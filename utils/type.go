package utils

type Nullable[T any] interface {
	~*T | interface{}
}
