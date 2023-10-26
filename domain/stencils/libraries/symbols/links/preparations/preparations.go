package preparations

type preparations struct {
	list []Preparation
}

func createPreparations(
	list []Preparation,
) Preparations {
	out := preparations{
		list: list,
	}

	return &out
}

// List returns the preparations
func (obj *preparations) List() []Preparation {
	return obj.list
}
