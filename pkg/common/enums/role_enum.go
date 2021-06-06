package enums

type Role string

const (
	USER        Role = "USER"
	ADMIN       Role = "ADMIN"
	ORG_ADMIN   Role = "ORG_ADMIN"
	SUPER_ADMIN Role = "SUPER_ADMIN"
)

func (role Role) IsSuperAdmin() bool {
	return role == SUPER_ADMIN
}

func (role Role) IsOrgAdmin() bool {
	switch role {
	case ORG_ADMIN:
	case SUPER_ADMIN:
		return true
	}

	return false
}

func (role Role) IsAdmin() bool {
	validRoles := []Role{ORG_ADMIN, SUPER_ADMIN, ADMIN}

	for i := 0; i < len(validRoles); i++ {
		if validRoles[i] == role {
			return true
		}
	}

	return false
}

func NewRoleFromString(role string) Role {

	validRoles := []Role{USER, ORG_ADMIN, SUPER_ADMIN, ADMIN}

	for i := 0; i < len(validRoles); i++ {
		if validRoles[i] == Role(role) {
			return Role(role)

		}
	}

	panic("invalid role provided, valid roles are provided")

}
