package utils

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func TestValidatorInvalidInput(t *testing.T) {

	user := User{}

	status, error := Validate(user)

	assert.False(t, status)
	assert.IsType(t, error, &validator.ValidationErrors{})

}

func TestValidatorValidInput(t *testing.T) {

	address := Address{
		Street: "Aloba Estate, Orogun, UI road",
		City:   "Ibadan",
		Planet: "Earth",
		Phone:  "08122222222",
	}

	status, error := Validate(address)

	assert.True(t, status)
	assert.Nil(t, error)

}
