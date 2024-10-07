package references

import "github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links/references/externals"

type reference struct {
	internal string
	external externals.External
}

func createReferenceWithInternal(
	internal string,
) Reference {
	return createReferenceInternally(internal, nil)
}

func createReferenceWithExternal(
	external externals.External,
) Reference {
	return createReferenceInternally("", external)
}

func createReferenceInternally(
	internal string,
	external externals.External,
) Reference {
	out := reference{
		internal: internal,
		external: external,
	}

	return &out
}

// IsInternal returns true if there is an internal, false otherwise
func (obj *reference) IsInternal() bool {
	return obj.internal != ""
}

// Internal returns the internal, if any
func (obj *reference) Internal() string {
	return obj.internal
}

// IsExternal returns true if there is an external, false otherwise
func (obj *reference) IsExternal() bool {
	return obj.external != nil
}

// External returns the external, if any
func (obj *reference) External() externals.External {
	return obj.external
}
