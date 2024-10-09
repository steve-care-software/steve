package assignments

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
)

// Assignment represents a query assignment
type Assignment interface {
	External() externals.External
	Variable() string
}
