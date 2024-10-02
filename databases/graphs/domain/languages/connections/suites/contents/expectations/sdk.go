package expectations

import "github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links/references"

// Expectation represents a suite expectation
type Expectation interface {
	References() references.Reference
	IsFail() bool
}
