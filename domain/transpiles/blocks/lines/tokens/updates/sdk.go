package updates

import "github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens/pointers"

// Update represents an update
type Update interface {
	Origin() pointers.Pointer
	Target() pointers.Pointer
}
