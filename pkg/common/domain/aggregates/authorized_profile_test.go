package aggregates

import (
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"suxenia-finance/pkg/common/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizeProfileEmail(t *testing.T) {

	authorizeProfile := AuthorizeProfile{}

	email := authorizeProfile.GetEmail()

	assert.IsType(t, objects.Email{}, email)

	newEmail := objects.NewEmail(utils.StrToPr("tayoadekunle@suxenia.com"))

	authorizeProfile.SetEmail(newEmail)

	assert.Equal(t, authorizeProfile.GetEmail(), newEmail)

}

func TestAuthorizeProfileFullName(t *testing.T) {

	authorizeProfile := AuthorizeProfile{}

	fullName, ok := authorizeProfile.GetFullName()

	assert.False(t, ok)

	assert.Nil(t, fullName)

	err := authorizeProfile.SetFullName("")

	assert.Error(t, err)

	name := "Tayo Adekunle"

	setError := authorizeProfile.SetFullName(name)

	assert.Nil(t, setError)

	fullName, ok = authorizeProfile.GetFullName()

	assert.True(t, ok)

	assert.Equal(t, fullName, name)

}

func TestAuthorizeProfileId(t *testing.T) {

	authorizeProfile := AuthorizeProfile{}

	assert.False(t, authorizeProfile.HasProfileId())

	profileId, ok := authorizeProfile.GetProfileId()

	assert.Nil(t, profileId)

	assert.False(t, ok)

	err := authorizeProfile.SetProfileId("")

	assert.Error(t, err)

	id := "suxenia-profile-id"

	err = authorizeProfile.SetProfileId(id)

	assert.Nil(t, err)

	profileId, ok = authorizeProfile.GetProfileId()

	assert.Equal(t, profileId, id)
	assert.True(t, ok)

}

func TestAuthorizeProfilePermissions(t *testing.T) {

	authorizeProfile := AuthorizeProfile{}

	permissions := authorizeProfile.GetPermissions()

	assert.IsType(t, permissions, objects.Permissions{})

	readPermission := objects.NewPermissionFromStrings([]string{"READ"})

	err := authorizeProfile.SetPermissions(readPermission)

	assert.Nil(t, err)

	assert.Equal(t, authorizeProfile.GetPermissions(), readPermission)

}

func TestAuthorizeOrgId(t *testing.T) {

	authorizeProfile := AuthorizeProfile{}

	assert.False(t, authorizeProfile.HasOrgId())

	orgId, ok := authorizeProfile.GetOrgId()

	assert.Nil(t, orgId)

	assert.False(t, ok)

	err := authorizeProfile.SetOrgId("")

	assert.Error(t, err)

	id := "suxenia-org-id"

	err = authorizeProfile.SetOrgId(id)

	assert.Nil(t, err)

	orgId, ok = authorizeProfile.GetOrgId()

	assert.Equal(t, orgId, id)

	assert.True(t, ok)

}
