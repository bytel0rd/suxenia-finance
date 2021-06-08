package structs

import (
	"suxenia-finance/pkg/common/enums"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthProfileIsValid(t *testing.T) {

	role := enums.ORG_ADMIN

	authProfile := AuthProfile{
		Email:       nil,
		FullName:    new(string),
		Id:          new(string),
		Permissions: &[]string{},
		Role:        &role,
		OrgId:       new(string),
	}

	ok := authProfile.IsValid()

	assert.False(t, ok)

}
