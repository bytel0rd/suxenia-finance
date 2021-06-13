package structs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIException(t *testing.T) {

	message := "Internal Server Error Exception"

	exception := NewAPIException(errors.New(message), 500)

	assert.Equal(t, exception.GetMessage(), message)

	assert.Equal(t, exception.GetStatusCode(), int(500))

	error := errors.New("UnTested Error Yet")

	errorCode := 503

	exception = NewAPIException(error, errorCode)

	assert.Equal(t, exception.GetMessage(), error.Error())

	assert.Equal(t, exception.GetStatusCode(), errorCode)

}

func TestNewAuthorizedException(t *testing.T) {
	exception := NewUnAuthorizedException(nil)

	assert.Equal(t, exception.GetStatusCode(), int(401))
}

func TestNewExceptionFromString(t *testing.T) {
	exception := NewAPIExceptionFromString("testing", 500)

	assert.Equal(t, exception.GetStatusCode(), int(500))
}

func TestNewBadRequestException(t *testing.T) {
	exception := NewBadRequestException(nil)

	assert.Equal(t, exception.GetStatusCode(), int(400))
}
