package headers

import "github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/headers/names"

type header struct {
	name    names.Name
	reverse names.Name
}

func createHeader(
	name names.Name,
) Header {
	return createHeaderInternally(name, nil)
}

func createHeaderWithReverse(
	name names.Name,
	reverse names.Name,
) Header {
	return createHeaderInternally(name, reverse)
}

func createHeaderInternally(
	name names.Name,
	reverse names.Name,
) Header {
	out := header{
		name:    name,
		reverse: reverse,
	}

	return &out
}

// Name returns the name
func (obj *header) Name() names.Name {
	return obj.name
}

// HasReverse returns true if reverse, false otherwise
func (obj *header) HasReverse() bool {
	return obj.reverse != nil
}

// Reverse returns the reverse, if any
func (obj *header) Reverse() names.Name {
	return obj.reverse
}
