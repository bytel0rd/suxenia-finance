package aggregates

import (
	"errors"
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"suxenia-finance/pkg/common/enums"
	"suxenia-finance/pkg/common/utils"
)

type AuthorizeProfile struct {
	email       objects.Email
	fullName    string
	id          string
	permissions objects.Permissions
	role        enums.Role
	orgId       string
}

func (profile *AuthorizeProfile) GetEmail() objects.Email {
	return profile.email
}

func (profile *AuthorizeProfile) SetEmail(email objects.Email) error {

	profile.email = email

	return nil
}

func (profile *AuthorizeProfile) GetFullName() (*string, bool) {

	if utils.IsValidString(profile.fullName) {
		return &profile.fullName, true
	}

	return nil, false
}

func (profile *AuthorizeProfile) SetFullName(name string) error {

	if name != "" {

		profile.fullName = name

		return nil
	}

	return errors.New("invalid full name provided")
}

func (profile *AuthorizeProfile) HasProfileId() bool {
	return utils.IsValidString(profile.id)
}

func (profile *AuthorizeProfile) GetProfileId() (*string, bool) {

	if profile.HasProfileId() {
		return &profile.id, true
	}

	return nil, false
}

func (profile *AuthorizeProfile) SetProfileId(profileId string) error {

	if utils.IsValidString(profileId) {
		profile.id = profileId
		return nil
	}

	return errors.New("invalid authorization profile id provided")
}

func (profile *AuthorizeProfile) GetPermissions() objects.Permissions {

	return profile.permissions
}

func (profile *AuthorizeProfile) SetPermissions(permissions objects.Permissions) error {

	profile.permissions = permissions

	return nil
}

func (profile *AuthorizeProfile) HasOrgId() bool {

	return utils.IsValidString(profile.orgId)
}

func (profile *AuthorizeProfile) GetOrgId() (*string, bool) {

	if profile.HasOrgId() {
		return &profile.orgId, true
	}

	return nil, false
}

func (profile *AuthorizeProfile) SetOrgId(orgId string) error {

	if utils.IsValidString(orgId) {
		profile.orgId = orgId
		return nil
	}

	return errors.New("Invalid organization Id provided")
}

func (profile *AuthorizeProfile) GetRole() enums.Role {

	return profile.role
}

func (profile *AuthorizeProfile) SetRole(role enums.Role) error {

	profile.role = role

	return nil

}

func NewAuthorizedProfile() AuthorizeProfile {
	return AuthorizeProfile{}
}
