package suites

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/suites/expectations"
)

type suite struct {
	name        string
	reference   references.Reference
	expectation expectations.Expectation
}

func createSuite(
	name string,
	reference references.Reference,
	expectation expectations.Expectation,
) Suite {
	out := suite{
		name:        name,
		reference:   reference,
		expectation: expectation,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Reference returns the reference
func (obj *suite) Reference() references.Reference {
	return obj.reference
}

// Expectation returns the expectation
func (obj *suite) Expectation() expectations.Expectation {
	return obj.expectation
}
