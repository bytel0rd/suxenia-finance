package objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionLength(t *testing.T) {

	inputPerms := []string{"READ", "WRITE"}

	perms := NewPermissionFromStrings(inputPerms)

	assert.Equal(t, perms.Length(), 2)

}

func TestPermissionAdd(t *testing.T) {

	inputPerms := []string{"READ", "WRITE"}

	perms := NewPermissionFromStrings(inputPerms)

	perms.Add("DELETE")

	assert.Equal(t, perms.Length(), 3)

}

func TestPermissionInclude(t *testing.T) {

	inputPerms := []string{"READ", "WRITE", "DELETE"}

	perms := NewPermissionFromStrings(inputPerms)

	ok := perms.Include("WRITE")

	assert.True(t, ok)

	ok = perms.Include("UPDATE")

	assert.False(t, ok)

}

func TestPermissionRemove(t *testing.T) {

	inputPerms := []string{"READ", "WRITE", "DELETE"}

	perms := NewPermissionFromStrings(inputPerms)

	perms.Remove("READ")

	assert.Equal(t, perms.Length(), 2)

}
