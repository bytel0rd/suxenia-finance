package utils

import (
	"database/sql"
	"testing"

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

	typeValue := []ValidatedFieldError{}

	assert.False(t, status)
	assert.IsType(t, error, &typeValue)

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

type Person struct {
	Name    sql.NullString `validate:"required"`
	PhoneNo sql.NullString `validate:"required,len=11"`
}

func TestValidatorSqlInValidInput(t *testing.T) {

	person := Person{
		Name:    sql.NullString{},
		PhoneNo: sql.NullString{},
	}

	status, error := Validate(person)

	typeValue := []ValidatedFieldError{}

	assert.False(t, status)
	assert.IsType(t, error, &typeValue)

}

func TestValidatorSqlValidInput(t *testing.T) {

	person := Person{
		Name:    sql.NullString{Valid: true, String: "Tayo Adekunle"},
		PhoneNo: sql.NullString{Valid: true, String: "08149464299"},
	}

	status, error := Validate(person)

	assert.True(t, status)
	assert.Nil(t, error)

}
