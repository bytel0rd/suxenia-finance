package objects

import "errors"

type Email struct {
	value *string
}

func (e Email) IsValid() bool {
	return emailValidator(e.value)
}

func (e Email) GetAddress() (*string, bool) {

	if e.IsValid() {
		return e.value, true
	}

	return nil, false
}

func (e Email) SetAddress(email *string) error {

	if email != nil && emailValidator(email) {
		e.value = email
		return nil
	}

	return errors.New("invalid email address provided")
}

func NewEmail(email *string) Email {
	return Email{
		value: email,
	}
}

func emailValidator(email *string) bool {
	return true
}
