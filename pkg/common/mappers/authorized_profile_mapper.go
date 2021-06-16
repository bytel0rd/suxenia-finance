package mappers

import (
	"suxenia-finance/pkg/common/domain/aggregates"
	objects "suxenia-finance/pkg/common/domain/valueobjects"
	"suxenia-finance/pkg/common/enums"
	"suxenia-finance/pkg/common/structs"
	"suxenia-finance/pkg/common/utils"
)

func NewAuthorizedProfileFromAuthProfile(profile structs.AuthProfile) (*aggregates.AuthorizeProfile, *structs.APIException) {

	if err := profile.Validate(); err != nil {

		apiError := structs.NewUnAuthorizedException(err)

		return nil, &apiError

	}

	authorizedProfile := aggregates.AuthorizeProfile{}

	authorizedProfile.SetProfileId(profile.ID)
	authorizedProfile.SetEmail(objects.NewEmail(utils.StrToPr(profile.Email)))
	authorizedProfile.SetFullName(profile.FullName)

	if profile.OrgID != nil {

		authorizedProfile.SetOrgId(*profile.OrgID)
	}

	authorizedProfile.SetPermissions(objects.NewPermissionFromStrings(profile.Permissions))
	authorizedProfile.SetRole(enums.NewRoleFromString(profile.Role))

	return &authorizedProfile, nil

}
