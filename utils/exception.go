package utils

import "errors"

type appException struct {
	Status  int
	Message error
}

type ApiException interface {
	appException | interface{}
}

func ApiInternalServerException(message string) ApiException {
	e := appException{}
	e.Status = 500
	e.Message = errors.New(message)
	return e
}

func ApiNotFoundException(message string) ApiException {
	e := appException{}
	e.Status = 404
	e.Message = errors.New(message)
	return e
}

func ApiConflictException(message string) ApiException {
	e := appException{}
	e.Status = 409
	e.Message = errors.New(message)
	return e
}
func ApiUnauthorizedException(message string) ApiException {
	e := appException{}
	e.Status = 401
	e.Message = errors.New(message)
	return e
}

func ApiBadRequestException(message string) ApiException {
	e := appException{}
	e.Status = 400
	e.Message = errors.New(message)
	return e
}

func ApiForbiddenException(message string) ApiException {
	e := appException{}
	e.Status = 403
	e.Message = errors.New(message)
	return e
}
