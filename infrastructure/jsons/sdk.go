package jsons

import (
	"github.com/steve-care-software/steve/applications/compilers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
	return_expectations "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/expectations"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/executions"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/origins"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/preparations"
	"github.com/steve-care-software/steve/domain/stencils/pointers"
	pointer_symbols "github.com/steve-care-software/steve/domain/stencils/pointers/symbols"
)

// NewJsonApplication creates a new json compiler application
func NewJsonApplication() compilers.Application {
	pRegistry := createRegistry()
	symbolsBuilder := symbols.NewBuilder()
	symbolBuilder := symbols.NewSymbolBuilder()
	pointersBuilder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	pointerSymbolBuilder := pointer_symbols.NewBuilder()
	linkBuilder := links.NewBuilder()
	linkExecutionBuilder := executions.NewBuilder()
	linkOriginsBuilder := origins.NewBuilder()
	linkOriginBuilder := origins.NewOriginBuilder()
	linkOriginDirectionBuilder := origins.NewDirectionBuilder()
	linkPreparationsBuilder := preparations.NewBuilder()
	linkPreparationBuilder := preparations.NewPreparationBuilder()
	linkConditionBuilder := preparations.NewConditionBuilder()
	layersBuilder := layers.NewBuilder()
	layerBuilder := layers.NewLayerBuilder()
	layerSuitesBuilder := layers.NewSuitesBuilder()
	layerSuiteBuilder := layers.NewSuiteBuilder()
	layerReturnBuilder := returns.NewBuilder()
	layerReturnExpectationBuilder := return_expectations.NewBuilder()
	layerReturnKindBuilder := kinds.NewBuilder()
	layerParamsBuilder := parameters.NewBuilder()
	layerParamBuilder := parameters.NewParameterBuilder()
	layerExecutionsBuilder := layers.NewExecutionsBuilder()
	layerExecutionBuilder := layers.NewExecutionBuilder()
	layerAssignmentBuilder := layers.NewAssignmentBuilder()
	layerConditionBuilder := layers.NewConditionBuilder()
	layerAssignableBuilder := layers.NewAssignableBuilder()
	layerReduceBuilder := reduces.NewBuilder()
	layerConstantValuesBuilder := layers.NewConstantValuesBuilder()
	layerConstantValueBuilder := layers.NewConstantValueBuilder()
	layerQueryBuilder := layers.NewQueryBuilder()
	layerLayerInputBuilder := layers.NewLayerInputBuilder()
	layerValueAssignmentsBuilder := layers.NewValueAssignmentsBuilder()
	layerValueAssignmentBuilder := layers.NewValueAssignmentBuilder()
	layerValueBuilder := layers.NewValueBuilder()
	return createCompilerApplication(
		pRegistry,
		symbolsBuilder,
		symbolBuilder,
		pointersBuilder,
		pointerBuilder,
		pointerSymbolBuilder,
		linkBuilder,
		linkExecutionBuilder,
		linkOriginsBuilder,
		linkOriginBuilder,
		linkOriginDirectionBuilder,
		linkPreparationsBuilder,
		linkPreparationBuilder,
		linkConditionBuilder,
		layersBuilder,
		layerBuilder,
		layerSuitesBuilder,
		layerSuiteBuilder,
		layerReturnBuilder,
		layerReturnExpectationBuilder,
		layerReturnKindBuilder,
		layerParamsBuilder,
		layerParamBuilder,
		layerExecutionsBuilder,
		layerExecutionBuilder,
		layerAssignmentBuilder,
		layerConditionBuilder,
		layerAssignableBuilder,
		layerReduceBuilder,
		layerConstantValuesBuilder,
		layerConstantValueBuilder,
		layerQueryBuilder,
		layerLayerInputBuilder,
		layerValueAssignmentsBuilder,
		layerValueAssignmentBuilder,
		layerValueBuilder,
	)
}
