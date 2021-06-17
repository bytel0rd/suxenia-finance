package structs

import (
	"errors"
	"suxenia-finance/pkg/common/utils"
)

type AuthProfile struct {
	Email       string   `json:"email" validate:"required,email"`
	FullName    string   `json:"fullName" validate:"required"`
	ID          string   `json:"id" validate:"required"`
	OrgID       *string  `json:"orgId,omitempty"`
	Permissions []string `json:"permissions" validate:"required"`
	Role        string   `json:"role" validate:"required"`
}

func (profile AuthProfile) Validate() error {

	_, fieldErrors := utils.Validate(profile)

	if fieldErrors != nil {
		errorCopy := *fieldErrors
		return errors.New(errorCopy[0].Message)
	}

	return nil
}

func (profile *AuthProfile) IsValid() bool {
	ok, _ := utils.Validate(profile)
	return ok
}
