package heads

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access"

type head struct {
	name    string
	version uint
	access  access.Access
}

func createHead(
	name string,
	version uint,
	access access.Access,
) Head {
	out := head{
		name:    name,
		version: version,
		access:  access,
	}

	return &out
}

// Name returns the name
func (obj *head) Name() string {
	return obj.name
}

// Version returns the version
func (obj *head) Version() uint {
	return obj.version
}

// Access returns the access
func (obj *head) Access() access.Access {
	return obj.access
}
