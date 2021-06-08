package structs

import "errors"

type StatusCode int64

type APIException struct {
	statusCode StatusCode
	error
}

func (e *APIException) GetStatusCode() StatusCode {
	return e.statusCode
}

func (e *APIException) GetMessage() string {
	return e.error.Error()
}

func (e *APIException) GetPtr() *APIException {
	return e
}

func NewStatusCode(code int64) StatusCode {
	return StatusCode(code)
}

func NewAPIException(error error, code *StatusCode) APIException {

	if code == nil {
		errorCode := NewStatusCode(500)
		code = &errorCode
	}

	exception := APIException{
		statusCode: *code,
		error:      error,
	}

	return exception
}

func NewUnAuthorizedException(error *error) APIException {

	if error == nil {
		errorMessage := errors.New("UnAuthorized Exception")
		error = &errorMessage
	}

	exception := APIException{
		statusCode: NewStatusCode(401),
		error:      *error,
	}

	return exception

}

func NewBadRequestException(error *error) APIException {

	if error == nil {
		errorMessage := errors.New("UnAuthorized Exception")
		error = &errorMessage
	}

	exception := APIException{
		statusCode: NewStatusCode(400),
		error:      *error,
	}

	return exception

}
