package suites

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites/expectations"
)

type suite struct {
	name        string
	link        links.Link
	expectation expectations.Expectations
}

func createSuite(
	name string,
	link links.Link,
	expectation expectations.Expectations,
) Suite {
	out := suite{
		name:        name,
		link:        link,
		expectation: expectation,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Link returns the link
func (obj *suite) Link() links.Link {
	return obj.link
}

// Expectations returns the expectation
func (obj *suite) Expectations() expectations.Expectations {
	return obj.expectation
}
