package objects

import (
	"suxenia-finance/pkg/common/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {

	email1 := NewEmail(nil)
	email2 := NewEmail(utils.StrToPr("a@"))
	email3 := NewEmail(utils.StrToPr("ade@gmail.com"))

	assert.False(t, email1.IsValid())
	assert.False(t, email2.IsValid())
	assert.True(t, email3.IsValid())

	value1, ok1 := email1.GetAddress()
	assert.False(t, ok1)
	assert.Nil(t, value1)

	value2, ok2 := email2.GetAddress()
	assert.False(t, ok2)
	assert.Nil(t, value2)

	value3, ok3 := email3.GetAddress()
	assert.True(t, ok3)
	assert.IsType(t, new(string), value3)

	err1 := email1.SetAddress(utils.StrToPr("ade@gmail.com"))
	assert.Nil(t, err1)

	err2 := email1.SetAddress(utils.StrToPr(""))
	assert.Error(t, err2)

}
