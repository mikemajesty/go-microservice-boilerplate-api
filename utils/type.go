package utils

type Nullable[T any] interface {
	~*T | interface{}
}

type PaginationType struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
type MongoListInput struct {
	Pagination PaginationType  `json:"pagination"`
	Sort       []MongoSortType `json:"sort"`
	Search     SearchType      `json:"search"`
}

type PostgresListInput struct {
	Pagination PaginationType `json:"pagination"`
	Sort       string         `json:"sort"`
	Search     SearchType     `json:"search"`
}
