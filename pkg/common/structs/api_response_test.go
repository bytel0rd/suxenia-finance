package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIResponse(t *testing.T) {

	data := "Success response"

	exception := NewAPIResponse(data, 200)

	assert.Equal(t, exception.GetData(), data)
	assert.Equal(t, exception.GetStatusCode(), 200)

}
