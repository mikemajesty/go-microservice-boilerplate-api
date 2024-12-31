package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ASC  = "asc"
	DESC = "desc"
)

type MongoSortType struct {
	Field string
	Order int
}

func MongoSort(controller *gin.Context) ([]MongoSortType, *AppException) {
	sort := controller.Query("sort")

	var sortList []MongoSortType

	if sort == "" {
		return []MongoSortType{{Field: "created_at", Order: 1}}, nil
	}

	for _, value := range strings.Split(sort, ",") {
		field := strings.Split(value, ":")

		if len(field) != 2 {
			return sortList, ApiBadRequestException("Invalid sort query")
		}

		sortField, err := MongoSortField(field[1])

		if err != nil {
			return sortList, ApiBadRequestException(err.Error())
		}

		sortList = append(sortList, MongoSortType{Field: field[0], Order: sortField.(int)})
	}

	return sortList, nil
}

func MongoSortField(field string) (Nullable[int], error) {
	if field == "" {
		return 1, nil
	}

	if field != ASC && field != DESC {
		return nil, errors.New("invalid sort field")
	}

	if field == ASC {
		return 1, nil
	}

	if field == DESC {
		return -1, nil
	}

	return nil, errors.New("invalid sort field")
}

func PostgresSort(controller *gin.Context) (Nullable[string], *AppException) {

	sort := controller.Query("sort")

	var sortList []string

	if sort == "" {
		return "created_at asc", nil
	}

	for _, value := range strings.Split(sort, ",") {
		field := strings.Split(value, ":")

		if len(field) != 2 {
			return nil, ApiBadRequestException("Invalid sort query")
		}

		sortList = append(sortList, field[0]+" "+field[1])
	}

	return strings.Join(sortList, ", "), nil
}
