package enums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleFromString(t *testing.T) {
	assert.Panics(t, func() {
		NewRoleFromString("TEST_ADMIN")
	})
}

func TestRoleIsSuperAdmin(t *testing.T) {
	assert.True(t, SUPER_ADMIN.IsSuperAdmin())
	assert.False(t, ORG_ADMIN.IsSuperAdmin())
}

func TestRoleIsOrgAdmin(t *testing.T) {
	v := ORG_ADMIN.IsOrgAdmin()
	fmt.Println(v)

	assert.True(t, SUPER_ADMIN.IsOrgAdmin())
	assert.True(t, ORG_ADMIN.IsOrgAdmin())
}

func TestRoleIsAdmin(t *testing.T) {
	assert.True(t, SUPER_ADMIN.IsAdmin())
	assert.True(t, ORG_ADMIN.IsAdmin())
	assert.True(t, ADMIN.IsAdmin())
	assert.False(t, USER.IsAdmin())
}
