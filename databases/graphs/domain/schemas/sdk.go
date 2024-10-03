package languages

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/headers"
)

// FetchGrammarInput fetches the grammar input
func FetchGrammarInput() []byte {
	return fetchGrammarInput()
}

// Adapter represents the schema adapter
type Adapter interface {
	ToLanguage(input []byte) (Schema, []byte, error)
}

// Schema represents the schema
type Schema interface {
	Header() headers.Header
	Points() []string
	Connections() connections.Connections
}
