package utils

type Nullable[T any] interface {
	~*T | interface{}
}

type PaginationType struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
type ListInput struct {
	Pagination PaginationType `json:"pagination"`
}
