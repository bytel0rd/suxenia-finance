package structs

import (
	"suxenia-finance/pkg/common/enums"
)

type AuthProfile struct {
	Email       *string
	FullName    *string
	Id          *string
	Permissions *[]string
	Role        *enums.Role
	OrgId       *string
}

func (profile *AuthProfile) IsValid() bool {
	return true
}
