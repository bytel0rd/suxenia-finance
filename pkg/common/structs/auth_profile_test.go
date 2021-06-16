package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthProfileIsValid(t *testing.T) {

	authProfile := AuthProfile{}

	ok := authProfile.IsValid()

	assert.False(t, ok)

	error := authProfile.Validate()
	assert.Error(t, error)

}
