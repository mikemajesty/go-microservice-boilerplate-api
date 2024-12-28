package utils

import "errors"

type AppException struct {
	Status  int
	Message error
}

type apiException interface {
	*AppException | interface{}
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
	e.Status = 409
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
