package objects

type Permissions struct {
	perms []string
}

func (p *Permissions) Length() int {
	return len(p.perms)
}

func (p *Permissions) Add(permission string) {
	p.perms = append(p.perms, permission)
}

func (p Permissions) Include(permission string) bool {

	includes := false

	for i := 0; i < len(p.perms); i++ {

		if p.perms[i] == permission {
			includes = true
		}

	}

	return includes
}

func (p Permissions) Remove(permission string) {

	filter := []string{}

	for i := 0; i < len(p.perms); i++ {

		if p.perms[i] != permission {
			filter = append(filter, p.perms[i])
		}

	}

	p.perms = filter

}

func NewPermission() Permissions {
	return Permissions{}
}

func NewPermissionFromStrings(perms *[]string) Permissions {

	if perms == nil {
		return NewPermission()
	}

	return Permissions{
		perms: *perms,
	}
}
