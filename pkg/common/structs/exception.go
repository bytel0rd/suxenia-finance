package structs

import (
	"errors"
	"net/http"
)

type APIException struct {
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"errors"`
	Message    string      `json:"errorMessage"`
}

func (e *APIException) GetStatusCode() int {
	return e.StatusCode
}

func (e *APIException) GetMessage() string {
	return e.Message
}

func (e *APIException) GetPtr() *APIException {
	return e
}

func NewAPIException(error error, code int) APIException {

	exception := APIException{
		StatusCode: code,
		Error:      error,
		Message:    error.Error(),
	}

	return exception
}

func NewAPIExceptionFromString(message string, code int) APIException {

	return NewAPIException(errors.New(message), code)
}

func NewUnAuthorizedException(error *error) APIException {

	if error == nil {
		errorMessage := errors.New("UnAuthorized Exception")
		error = &errorMessage
	}

	return NewAPIException(*error, http.StatusUnauthorized)

}

func NewBadRequestException(error *error) APIException {

	if error == nil {
		errorMessage := errors.New("Bad Exception")
		error = &errorMessage
	}

	return NewAPIException(*error, http.StatusBadRequest)

}
