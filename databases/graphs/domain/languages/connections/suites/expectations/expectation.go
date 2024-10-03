package expectations

import "github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links/references"

type expectation struct {
	references references.References
	isFail     bool
}

func createExpectation(
	references references.References,
	isFail bool,
) Expectation {
	out := expectation{
		references: references,
		isFail:     isFail,
	}

	return &out
}

// References returns the references
func (obj *expectation) References() references.References {
	return obj.references
}

// IsFail returns true if fail, false otherwise
func (obj *expectation) IsFail() bool {
	return obj.isFail
}
