package structs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbException(t *testing.T) {
	error := errors.New("DB Exception")
	dbException := NewDBException(error, false)

	assert.Equal(t, dbException.Error(), error.Error())
	assert.True(t, dbException.IsExecError)
	assert.False(t, dbException.isValidationError)
}
