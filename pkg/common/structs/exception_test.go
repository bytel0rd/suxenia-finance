package structs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIException(t *testing.T) {

	message := "Internal Server Error Exception"

	exception := NewAPIException(errors.New(message), nil)

	assert.Equal(t, exception.GetMessage(), message)

	assert.Equal(t, exception.GetStatusCode(), NewStatusCode(500))

	error := errors.New("UnTested Error Yet")

	errorCode := NewStatusCode(503)

	exception = NewAPIException(error, &errorCode)

	assert.Equal(t, exception.GetMessage(), error.Error())

	assert.Equal(t, exception.GetStatusCode(), errorCode)

}

func TestNewAuthorizedException(t *testing.T) {
	exception := NewUnAuthorizedException(nil)

	assert.Equal(t, exception.GetStatusCode(), NewStatusCode(401))
}

func TestNewBadRequestException(t *testing.T) {
	exception := NewBadRequestException(nil)

	assert.Equal(t, exception.GetStatusCode(), NewStatusCode(400))
}
