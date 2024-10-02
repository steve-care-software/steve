package languages

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections"
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/headers"
)

// FetchGrammarInput fetches the grammar input
func FetchGrammarInput() []byte {
	return fetchGrammarInput()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithGrammar(grammar []byte) AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents the language adapter
type Adapter interface {
	ToLanguage(input []byte) (Language, []byte, error)
}

// Language represents the language
type Language interface {
	Header() headers.Header
	Points() []string
	Connections() connections.Connections
}
