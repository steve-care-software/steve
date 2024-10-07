package scripts

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/permissions"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/writes"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections"
	connection_headers "github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/headers"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/headers/names"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/headers/names/cardinalities"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links/references/externals"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/suites"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/suites/expectations"
	applications_parser "github.com/steve-care-software/steve/parsers/applications"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/walkers/elements"
)

type adapter struct {
	parserAppBuilder        applications_parser.Builder
	grammar                 grammars.Grammar
	builder                 Builder
	headBuilder             heads.Builder
	accessBuilder           access.Builder
	permissionsBuilder      permissions.Builder
	permissionBuilder       permissions.PermissionBuilder
	writeBuilder            writes.Builder
	connectionsBuilder      connections.Builder
	connectionBuilder       connections.ConnectionBuilder
	suitesBuilder           suites.Builder
	suiteBuilder            suites.SuiteBuilder
	expectationsBuilder     expectations.Builder
	expectationBuilder      expectations.ExpectationBuilder
	linksBuilder            links.Builder
	linkBuilder             links.LinkBuilder
	referencesBuilder       references.Builder
	referenceBuilder        references.ReferenceBuilder
	externalBuilder         externals.Builder
	connectionHeaderBuilder connection_headers.Builder
	nameBuilder             names.Builder
	cardinalityBuilder      cardinalities.Builder
}

func createAdapter(
	parserAppBuilder applications_parser.Builder,
	grammar grammars.Grammar,
	builder Builder,
	headBuilder heads.Builder,
	accessBuilder access.Builder,
	permissionsBuilder permissions.Builder,
	permissionBuilder permissions.PermissionBuilder,
	writeBuilder writes.Builder,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	suitesBuilder suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	expectationsBuilder expectations.Builder,
	expectationBuilder expectations.ExpectationBuilder,
	linksBuilder links.Builder,
	linkBuilder links.LinkBuilder,
	referencesBuilder references.Builder,
	referenceBuilder references.ReferenceBuilder,
	externalBuilder externals.Builder,
	connectionHeaderBuilder connection_headers.Builder,
	nameBuilder names.Builder,
	cardinalityBuilder cardinalities.Builder,

) Adapter {
	out := adapter{
		parserAppBuilder:        parserAppBuilder,
		grammar:                 grammar,
		builder:                 builder,
		headBuilder:             headBuilder,
		accessBuilder:           accessBuilder,
		permissionsBuilder:      permissionsBuilder,
		permissionBuilder:       permissionBuilder,
		writeBuilder:            writeBuilder,
		connectionsBuilder:      connectionsBuilder,
		connectionBuilder:       connectionBuilder,
		suitesBuilder:           suitesBuilder,
		suiteBuilder:            suiteBuilder,
		expectationsBuilder:     expectationsBuilder,
		expectationBuilder:      expectationBuilder,
		linksBuilder:            linksBuilder,
		linkBuilder:             linkBuilder,
		referencesBuilder:       referencesBuilder,
		referenceBuilder:        referenceBuilder,
		externalBuilder:         externalBuilder,
		connectionHeaderBuilder: connectionHeaderBuilder,
		nameBuilder:             nameBuilder,
		cardinalityBuilder:      cardinalityBuilder,
	}

	return &out
}

// ToScript converts an input to script
func (app *adapter) ToScript(input []byte) (Script, []byte, error) {
	sequence := elements.Element{
		ElementFn: func(input any) (any, error) {
			return input, nil
		},
		TokenList: &elements.TokenList{
			MapFn: func(elementName string, mp map[string][]any) (any, error) {
				return nil, nil
			},
			List: map[string]elements.SelectedTokenList{},
		},
	}

	parserApp, err := app.parserAppBuilder.Create().WithElement(sequence).Now()
	if err != nil {
		return nil, nil, err
	}

	retIns, retRemaining, err := parserApp.Execute(input, app.grammar)
	if err != nil {
		return nil, nil, err
	}

	if casted, ok := retIns.(Script); ok {
		return casted, retRemaining, nil
	}

	return nil, nil, errors.New("the returned parser instance was was expected to contain a Script instance")
}
