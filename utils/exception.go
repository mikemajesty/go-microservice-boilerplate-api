package utils

type AppException struct {
	Status  int
	Message string
}

func (e *AppException) ApiInternalServerException(message string) AppException {
	e.Status = 500
	e.Message = message
	return *e
}

func (e *AppException) ApiNotFoundException(message string) AppException {
	e.Status = 404
	e.Message = message
	return *e
}

func (e *AppException) ApiConflictException(message string) AppException {
	e.Status = 409
	e.Message = message
	return *e
}
func (e *AppException) ApiUnauthorizedException(message string) AppException {
	e.Status = 401
	e.Message = message
	return *e
}

func (e *AppException) ApiBadRequestException(message string) AppException {
	e.Status = 400
	e.Message = message
	return *e
}

func (e *AppException) ApiForbiddenException(message string) AppException {
	e.Status = 403
	e.Message = message
	return *e
}

func (e *AppException) GetStatus() int {
	return e.Status
}

func (e *AppException) GetMessage() string {
	return e.Message
}
