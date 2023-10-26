package origins

type direction struct {
	next     Origin
	previous Origin
}

func createDirectionWithNext(
	next Origin,
) Direction {
	return createDirectionInternally(next, nil)
}

func createDirectionWithPrevious(
	previous Origin,
) Direction {
	return createDirectionInternally(nil, previous)
}

func createDirectionInternally(
	next Origin,
	previous Origin,
) Direction {
	out := direction{
		next:     next,
		previous: previous,
	}

	return &out
}

// IsNext returns true if there is a next, false otherwise
func (obj *direction) IsNext() bool {
	return obj.next != nil
}

// Next returns the next, if any
func (obj *direction) Next() Origin {
	return obj.next
}

// IsPrevious returns true if there is a previous, false otherwise
func (obj *direction) IsPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous, if any
func (obj *direction) Previous() Origin {
	return obj.previous
}
