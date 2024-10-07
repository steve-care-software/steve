package permissions

type permission struct {
	names        []string
	compensation float64
}

func createPermission(
	names []string,
) Permission {
	return createPermissionInternally(
		names,
		0.0,
	)
}

func createPermissionWithCompensation(
	names []string,
	compensation float64,
) Permission {
	return createPermissionInternally(
		names,
		compensation,
	)
}

func createPermissionInternally(
	names []string,
	compensation float64,
) Permission {
	out := permission{
		names:        names,
		compensation: compensation,
	}

	return &out
}

// Names returns the names
func (obj *permission) Names() []string {
	return obj.names
}

// HasCompensation returns true if there is a compensation, false otherwise
func (obj *permission) HasCompensation() bool {
	return obj.compensation > 0.0
}

// Compensation returns the compensation
func (obj *permission) Compensation() float64 {
	return obj.compensation
}
