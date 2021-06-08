package mappers

import (
	"suxenia-finance/pkg/common/domain/aggregates"
	"suxenia-finance/pkg/common/enums"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizedProfileFromAuthProfileMapper(t *testing.T) {

	role := enums.ORG_ADMIN

	profile := structs.AuthProfile{
		Email:       new(string),
		FullName:    new(string),
		Id:          new(string),
		Permissions: &[]string{},
		Role:        &role,
		OrgId:       new(string),
	}

	authorizedProfile, err := NewAuthorizedProfileFromAuthProfile(profile)

	assert.Nil(t, authorizedProfile)
	assert.IsType(t, *err, structs.APIException{})

	profile = structs.AuthProfile{
		Email:       utils.StrToPr("tayoadekunle@suxenia.com"),
		FullName:    utils.StrToPr("Tayo Adekunle"),
		Id:          utils.StrToPr("suxenia-profile-id"),
		Permissions: &[]string{},
		Role:        &role,
		OrgId:       utils.StrToPr("suxenia-orgoid"),
	}

	authorizedProfile, err = NewAuthorizedProfileFromAuthProfile(profile)

	assert.IsType(t, *authorizedProfile, aggregates.AuthorizeProfile{})
	assert.Nil(t, err)

}
