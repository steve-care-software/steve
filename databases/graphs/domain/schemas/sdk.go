package schemas

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/headers"
)

// NewAdapterFactory creates a new adapter factory
/*func NewAdapterFactory() AdapterFactory {
	astAdapter := asts.NewAdapter()
	parserAdapter := grammars.NewAdapter()
	queryAdapterFactory := queries.NewAdapterFactory()
	builder := NewBuilder()
	headerBuilder := headers.NewBuilder()
	connectionsBuilder := connections.NewBuilder()
	connectionBuilder := connections.NewConnectionBuilder()
	suitesBuilder := suites.NewBuilder()
	suiteBuilder := suites.NewSuiteBuilder()
	expectationsBuilder := expectations.NewBuilder()
	expectationBuilder := expectations.NewExpectationBuilder()
	linksBuilder := links.NewBuilder()
	linkBuilder := links.NewLinkBuilder()
	referencesBuilder := references.NewBuilder()
	referenceBuilder := references.NewReferenceBuilder()
	externalBuilder := externals.NewBuilder()
	connectionHeaderBuilder := connection_headers.NewBuilder()
	nameBuilder := names.NewBuilder()
	cardinalityBuilder := cardinalities.NewBuilder()
	grammarInput := fetchGrammarInput()
	return createAdapterFactory(
		astAdapter,
		parserAdapter,
		queryAdapterFactory,
		builder,
		headerBuilder,
		connectionsBuilder,
		connectionBuilder,
		suitesBuilder,
		suiteBuilder,
		expectationsBuilder,
		expectationBuilder,
		linksBuilder,
		linkBuilder,
		referencesBuilder,
		referenceBuilder,
		externalBuilder,
		connectionHeaderBuilder,
		nameBuilder,
		cardinalityBuilder,
		grammarInput,
	)
}*/

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// AdapterFactory represents an adapter factory
type AdapterFactory interface {
	Create() (Adapter, error)
}

// Adapter represents the schema adapter
type Adapter interface {
	ToSchema(input []byte) (Schema, []byte, error)
}

// Builder represents the schema builder
type Builder interface {
	Create() Builder
	WithHeader(header headers.Header) Builder
	WithPoints(points []string) Builder
	WithConnections(connections connections.Connections) Builder
	Now() (Schema, error)
}

// Schema represents the schema
type Schema interface {
	Header() headers.Header
	Points() []string
	Connections() connections.Connections
}
