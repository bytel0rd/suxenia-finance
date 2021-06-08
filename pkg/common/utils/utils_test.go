package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidString(t *testing.T) {
	emptyString := ""
	validString := "Random string"

	assert.False(t, IsValidString(nil))

	assert.False(t, IsValidString(&emptyString))
	assert.False(t, IsValidString(new(string)))

	assert.True(t, IsValidString(&validString))

}
