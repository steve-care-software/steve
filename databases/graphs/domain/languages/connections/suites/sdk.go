package suites

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/suites/contents"
)

// Suites represents suites
type Suites interface {
	List() []Suite
}

// Suite represents a suite
type Suite interface {
	Name() string
	Journey() references.Reference
	Content() contents.Content
}
