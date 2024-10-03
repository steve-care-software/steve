package headers

type header struct {
	version uint
	name    string
}

func createHeader(
	version uint,
	name string,
) Header {
	out := header{
		version: version,
		name:    name,
	}

	return &out
}

// Version returns the version
func (obj *header) Version() uint {
	return obj.version
}

// Name returns the name
func (obj *header) Name() string {
	return obj.name
}
