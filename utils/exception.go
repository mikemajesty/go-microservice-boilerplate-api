package utils

import (
	"errors"
)

type AppException struct {
	Status  int
	Message error
}

var mapMessage = map[int]Nullable[string]{
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	409: "Conflict",
	500: "Oops, an error occurred",
}

func getMessage(err *AppException) string {
	var message = mapMessage[err.GetStatus()]

	if message != nil {
		return message.(string)
	}

	return err.Message.Error()
}

func (e *AppException) Response(status int, traceId string) any {
	return struct {
		Message string `json:"message"`
		TraceID string `json:"trace_id"`
		Status  int    `json:"status"`
	}{
		Message: getMessage(e),
		Status:  status,
		TraceID: traceId,
	}
}

func (e *AppException) GetMessage() string {
	return e.Message.Error()
}

func (e *AppException) GetStatus() int {
	return e.Status
}

func ApiInternalServerException(message string) *AppException {
	e := AppException{}
	e.Status = 500
	e.Message = errors.New(message)
	return &e
}

func ApiNotFoundException(message string) *AppException {
	e := AppException{}
	e.Status = 404
	e.Message = errors.New(message)
	return &e
}

func ApiConflictException(message string) *AppException {
	e := AppException{}
	e.Status = 406
	e.Message = errors.New(message)
	return &e
}
func ApiUnauthorizedException(message string) *AppException {
	e := AppException{}
	e.Status = 401
	e.Message = errors.New(message)
	return &e
}

func ApiBadRequestException(message string) *AppException {
	e := AppException{}
	e.Status = 400
	e.Message = errors.New(message)
	return &e
}

func ApiForbiddenException(message string) *AppException {
	e := AppException{}
	e.Status = 403
	e.Message = errors.New(message)
	return &e
}
