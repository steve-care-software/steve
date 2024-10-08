package expectations

type expectations struct {
	list []Expectation
}

func createExpectations(
	list []Expectation,
) Expectations {
	out := expectations{
		list: list,
	}

	return &out
}

// List returns the list of expectation
func (obj *expectations) List() []Expectation {
	return obj.list
}
