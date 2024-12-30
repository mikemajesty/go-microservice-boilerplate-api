package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(controller *gin.Context) PaginationType {
	var pagination = PaginationType{Page: 1, Limit: 10}
	pageNumber, erroPage := strconv.Atoi(controller.Query("page"))

	if erroPage == nil {
		pagination.Page = pageNumber
	}

	if pagination.Page < 1 {
		pagination.Page = 1
	}

	limit, erroLimit := strconv.Atoi(controller.Query("limit"))

	if erroLimit == nil {
		pagination.Limit = limit
	}

	if pagination.Limit < 1 {
		pagination.Limit = 10
	}

	if limit > 100 {
		pagination.Limit = 100
	}

	return pagination
}
