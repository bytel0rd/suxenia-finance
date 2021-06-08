package mappers

import (
	"suxenia-finance/pkg/common/domain/aggregates"
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"suxenia-finance/pkg/common/enums"
	"suxenia-finance/pkg/common/structs"
)

func NewAuthorizedProfileFromAuthProfile(profile structs.AuthProfile) (*aggregates.AuthorizeProfile, *structs.APIException) {

	if err := profile.Validate(); err != nil {

		apiError := structs.NewUnAuthorizedException(&err)

		return nil, &apiError

	}

	authorizedProfile := aggregates.AuthorizeProfile{}

	authorizedProfile.SetEmail(objects.NewEmail(profile.Email))
	authorizedProfile.SetFullName(*profile.FullName)
	authorizedProfile.SetOrgId(*profile.OrgId)
	authorizedProfile.SetPermissions(objects.NewPermissionFromStrings(profile.Permissions))
	authorizedProfile.SetRole(enums.NewRoleFromString(string(*profile.Role)))

	return &authorizedProfile, nil

}
