package schemas

import (
	"errors"
	"fmt"
	"strconv"

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
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

type adapter struct {
	astAdapter              asts.Adapter
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
	grammar                 grammars.Grammar
}

func createAdapter(
	astAdapter asts.Adapter,
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
	grammar grammars.Grammar,
) Adapter {
	out := adapter{
		astAdapter:              astAdapter,
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
		grammar:                 grammar,
	}

	return &out
}

// ToSchema converts an input to a Schema instance
func (app *adapter) ToSchema(input []byte) (Schema, []byte, error) {
	ast, retRemaining, err := app.astAdapter.ToAST(app.grammar, input)
	if err != nil {
		return nil, nil, err
	}

	root := ast.Root()
	retSchema, err := app.schema(root)
	if err != nil {
		return nil, nil, err
	}

	return retSchema, retRemaining, nil
}

func (app *adapter) schema(element instructions.Element) (Schema, error) {
	tokens, err := app.elementToTokens(element)
	if err != nil {
		return nil, err
	}

	retHeadToken, err := tokens.Fetch("head", 0)
	if err != nil {
		return nil, err
	}

	retPointsToken, err := tokens.Fetch("instructionPoints", 0)
	if err != nil {
		return nil, err
	}

	retConnectionsToken, err := tokens.Fetch("connectionBlocks", 0)
	if err != nil {
		return nil, err
	}

	retHead, err := app.head(retHeadToken)
	if err != nil {
		return nil, err
	}

	retPoints, err := app.points(retPointsToken)
	if err != nil {
		return nil, err
	}

	retConnections, err := app.connections(retConnectionsToken)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithHeader(retHead).
		WithPoints(retPoints).
		WithConnections(retConnections).
		Now()
}

func (app *adapter) head(token instructions.Token) (headers.Header, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return nil, err
	}

	retVersionToken, err := tokens.Fetch("versionInstruction", 0)
	if err != nil {
		return nil, err
	}

	versionNumber, err := app.versionInstruction(retVersionToken)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nhead: %d\n", versionNumber)
	return nil, nil
}

func (app *adapter) versionInstruction(token instructions.Token) (int, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return 0, err
	}

	retNumberToken, err := tokens.Fetch("numbers", 0)
	if err != nil {
		return 0, err
	}

	retBytes := retNumberToken.Value()
	return strconv.Atoi(string(retBytes))
}

func (app *adapter) points(token instructions.Token) ([]string, error) {
	return nil, nil
}

func (app *adapter) connections(token instructions.Token) (connections.Connections, error) {
	return nil, nil
}

func (app *adapter) tokenToFirstInstructionTokens(token instructions.Token) (instructions.Tokens, error) {
	retElement, err := token.Elements().Fetch(0)
	if err != nil {
		return nil, err
	}

	return app.elementToTokens(retElement)
}

func (app *adapter) elementToTokens(element instructions.Element) (instructions.Tokens, error) {
	if element.IsConstant() {
		str := fmt.Sprintf("the element (name: %s) was expected to contain an Instruction", element.Name())
		return nil, errors.New(str)
	}

	return element.Instruction().Tokens(), nil
}
