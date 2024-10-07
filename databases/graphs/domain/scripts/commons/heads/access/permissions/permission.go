package permissions

type permission struct {
	name         string
	compensation float64
}

func createPermission(
	name string,
	compensation float64,
) Permission {
	out := permission{
		name:         name,
		compensation: compensation,
	}

	return &out
}

// Name returns the name
func (obj *permission) Name() string {
	return obj.name
}

// Compensation returns the compensation
func (obj *permission) Compensation() float64 {
	return obj.compensation
}
