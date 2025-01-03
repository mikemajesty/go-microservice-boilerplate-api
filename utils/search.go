package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type SearchType struct {
	Field string
	Value string
}

func CreateSearch(controller *gin.Context) (SearchType, *AppException) {
	search := controller.Query("search")

	if search == "" {
		return SearchType{}, nil
	}

	var searchData SearchType

	for _, value := range strings.Split(search, ",") {
		field := strings.Split(value, ":")

		if len(field) != 2 {
			return searchData, ApiBadRequestException("Invalid search query")
		}

		searchData = SearchType{Field: field[0], Value: field[1]}
	}

	return searchData, nil
}
