package contents

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/suites/contents/expectations"
)

// Content represents the content
type Content interface {
	IsOptimal() bool
	Optimal() references.References
	IsExpectation() bool
	Expectation() expectations.Expectation
}
