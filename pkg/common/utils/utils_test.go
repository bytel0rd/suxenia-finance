package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidString(t *testing.T) {
	emptyString := ""
	validString := "Random string"

	assert.False(t, IsValidString(""))
	assert.True(t, IsValidString("Tayo Adekunle"))

	assert.False(t, IsValidStringPointer(&emptyString))
	assert.False(t, IsValidStringPointer(new(string)))

	assert.True(t, IsValidStringPointer(&validString))

}
