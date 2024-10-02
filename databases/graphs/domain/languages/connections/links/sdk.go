package links

import "github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links/references"

// Links represents links
type Links interface {
	List() []Link
}

// Link represents a link
type Link interface {
	Origin() references.Reference
	Target() references.Reference
}
