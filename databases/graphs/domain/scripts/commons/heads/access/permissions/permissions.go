package permissions

type permissions struct {
	list []Permission
}

func createPermissions(
	list []Permission,
) Permissions {
	out := permissions{
		list: list,
	}

	return &out
}

// List returns the list of permission
func (obj *permissions) List() []Permission {
	return obj.list
}
