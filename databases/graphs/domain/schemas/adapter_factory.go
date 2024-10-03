package schemas

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections"
	connection_headers "github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers/names"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers/names/cardinalities"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references/externals"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites/expectations"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/headers"
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/queries"
)

type adapterFactory struct {
	astAdapter              asts.Adapter
	parserAdapter           grammars.Adapter
	queryAdapterFactory     queries.AdapterFactory
	builder                 Builder
	headerBuilder           headers.Builder
	connectionsBuilder      connections.Builder
	connectionBuilder       connections.ConnectionBuilder
	suitesBuilder           suites.Builder
	suiteBuilder            suites.SuiteBuilder
	expectationBuilder      expectations.Builder
	linksBuilder            links.Builder
	linkBuilder             links.LinkBuilder
	referencesBuilder       references.Builder
	referenceBuilder        references.ReferenceBuilder
	externalBuilder         externals.Builder
	connectionHeaderBuilder connection_headers.Builder
	nameBuilder             names.Builder
	cardinalityBuilder      cardinalities.Builder
	input                   []byte
}

func createAdapterFactory(
	astAdapter asts.Adapter,
	parserAdapter grammars.Adapter,
	queryAdapterFactory queries.AdapterFactory,
	builder Builder,
	headerBuilder headers.Builder,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	suitesBuilder suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	expectationBuilder expectations.Builder,
	linksBuilder links.Builder,
	linkBuilder links.LinkBuilder,
	referencesBuilder references.Builder,
	referenceBuilder references.ReferenceBuilder,
	externalBuilder externals.Builder,
	connectionHeaderBuilder connection_headers.Builder,
	nameBuilder names.Builder,
	cardinalityBuilder cardinalities.Builder,
	input []byte,
) AdapterFactory {
	out := adapterFactory{
		astAdapter:              astAdapter,
		parserAdapter:           parserAdapter,
		queryAdapterFactory:     queryAdapterFactory,
		builder:                 builder,
		headerBuilder:           headerBuilder,
		connectionsBuilder:      connectionsBuilder,
		connectionBuilder:       connectionBuilder,
		suitesBuilder:           suitesBuilder,
		suiteBuilder:            suiteBuilder,
		expectationBuilder:      expectationBuilder,
		linksBuilder:            linksBuilder,
		linkBuilder:             linkBuilder,
		referencesBuilder:       referencesBuilder,
		referenceBuilder:        referenceBuilder,
		externalBuilder:         externalBuilder,
		connectionHeaderBuilder: connectionHeaderBuilder,
		nameBuilder:             nameBuilder,
		cardinalityBuilder:      cardinalityBuilder,
		input:                   input,
	}

	return &out
}

// Create creates the adapter
func (app *adapterFactory) Create() (Adapter, error) {
	grammar, _, err := app.parserAdapter.ToGrammar(app.input)
	if err != nil {
		return nil, err
	}

	queryAdapter, err := app.queryAdapterFactory.Create()
	if err != nil {
		return nil, err
	}

	return createAdapter(
		app.astAdapter,
		queryAdapter,
		app.builder,
		app.headerBuilder,
		app.connectionsBuilder,
		app.connectionBuilder,
		app.suitesBuilder,
		app.suiteBuilder,
		app.expectationBuilder,
		app.linksBuilder,
		app.linkBuilder,
		app.referencesBuilder,
		app.referenceBuilder,
		app.externalBuilder,
		app.connectionHeaderBuilder,
		app.nameBuilder,
		app.cardinalityBuilder,
		grammar,
	), nil
}
