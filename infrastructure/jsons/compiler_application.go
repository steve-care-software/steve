package jsons

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/steve/applications/compilers"
	"github.com/steve-care-software/steve/domain/pointers"
	pointer_symbols "github.com/steve-care-software/steve/domain/pointers/symbols"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
	return_expectations "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/expectations"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/executions"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/origins"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/preparations"
	"github.com/steve-care-software/steve/infrastructure/jsons/structs/layers/assignables/administrators"
	struct_values "github.com/steve-care-software/steve/infrastructure/jsons/structs/values"
)

type compilerApplication struct {
	pRegistry                     *registry
	symbolsBuilder                symbols.Builder
	symbolBuilder                 symbols.SymbolBuilder
	pointersBuilder               pointers.Builder
	pointerBuilder                pointers.PointerBuilder
	pointerSymbolBuilder          pointer_symbols.Builder
	linkBuilder                   links.Builder
	linkExecutionBuilder          executions.Builder
	linkOriginsBuilder            origins.Builder
	linkOriginBuilder             origins.OriginBuilder
	linkOriginDirectionBuilder    origins.DirectionBuilder
	linkPreparationsBuilder       preparations.Builder
	linkPreparationBuilder        preparations.PreparationBuilder
	linkConditionBuilder          preparations.ConditionBuilder
	layersBuilder                 layers.Builder
	layerBuilder                  layers.LayerBuilder
	layerSuitesBuilder            layers.SuitesBuilder
	layerSuiteBuilder             layers.SuiteBuilder
	layerReturnBuilder            returns.Builder
	layerReturnExpectationBuilder return_expectations.Builder
	layerReturnKindBuilder        kinds.Builder
	layerParamsBuilder            parameters.Builder
	layerParamBuilder             parameters.ParameterBuilder
	layerExecutionsBuilder        layers.ExecutionsBuilder
	layerExecutionBuilder         layers.ExecutionBuilder
	layerAssignmentBuilder        layers.AssignmentBuilder
	layerConditionBuilder         layers.ConditionBuilder
	layerAssignableBuilder        layers.AssignableBuilder
	layerReduceBuilder            reduces.Builder
	layerConstantValuesBuilder    constantvalues.ConstantValuesBuilder
	layerConstantValueBuilder     constantvalues.ConstantValueBuilder
	layerQueryBuilder             layers.QueryBuilder
	layerLayerInputBuilder        layers.LayerInputBuilder
	layerValueAssignmentsBuilder  layers.ValueAssignmentsBuilder
	layerValueAssignmentBuilder   layers.ValueAssignmentBuilder
	layerValueBuilder             layers.ValueBuilder
}

func createCompilerApplication(
	pRegistry *registry,
	symbolsBuilder symbols.Builder,
	symbolBuilder symbols.SymbolBuilder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	pointerSymbolBuilder pointer_symbols.Builder,
	linkBuilder links.Builder,
	linkExecutionBuilder executions.Builder,
	linkOriginsBuilder origins.Builder,
	linkOriginBuilder origins.OriginBuilder,
	linkOriginDirectionBuilder origins.DirectionBuilder,
	linkPreparationsBuilder preparations.Builder,
	linkPreparationBuilder preparations.PreparationBuilder,
	linkConditionBuilder preparations.ConditionBuilder,
	layersBuilder layers.Builder,
	layerBuilder layers.LayerBuilder,
	layerSuitesBuilder layers.SuitesBuilder,
	layerSuiteBuilder layers.SuiteBuilder,
	layerReturnBuilder returns.Builder,
	layerReturnExpectationBuilder return_expectations.Builder,
	layerReturnKindBuilder kinds.Builder,
	layerParamsBuilder parameters.Builder,
	layerParamBuilder parameters.ParameterBuilder,
	layerExecutionsBuilder layers.ExecutionsBuilder,
	layerExecutionBuilder layers.ExecutionBuilder,
	layerAssignmentBuilder layers.AssignmentBuilder,
	layerConditionBuilder layers.ConditionBuilder,
	layerAssignableBuilder layers.AssignableBuilder,
	layerReduceBuilder reduces.Builder,
	layerConstantValuesBuilder constantvalues.ConstantValuesBuilder,
	layerConstantValueBuilder constantvalues.ConstantValueBuilder,
	layerQueryBuilder layers.QueryBuilder,
	layerLayerInputBuilder layers.LayerInputBuilder,
	layerValueAssignmentsBuilder layers.ValueAssignmentsBuilder,
	layerValueAssignmentBuilder layers.ValueAssignmentBuilder,
	layerValueBuilder layers.ValueBuilder,
) compilers.Application {
	out := compilerApplication{
		pRegistry:                     pRegistry,
		symbolsBuilder:                symbolsBuilder,
		symbolBuilder:                 symbolBuilder,
		pointersBuilder:               pointersBuilder,
		pointerBuilder:                pointerBuilder,
		pointerSymbolBuilder:          pointerSymbolBuilder,
		linkBuilder:                   linkBuilder,
		linkExecutionBuilder:          linkExecutionBuilder,
		linkOriginsBuilder:            linkOriginsBuilder,
		linkOriginBuilder:             linkOriginBuilder,
		linkOriginDirectionBuilder:    linkOriginDirectionBuilder,
		linkPreparationsBuilder:       linkPreparationsBuilder,
		linkPreparationBuilder:        linkPreparationBuilder,
		linkConditionBuilder:          linkConditionBuilder,
		layersBuilder:                 layersBuilder,
		layerBuilder:                  layerBuilder,
		layerSuitesBuilder:            layerSuitesBuilder,
		layerSuiteBuilder:             layerSuiteBuilder,
		layerReturnBuilder:            layerReturnBuilder,
		layerReturnExpectationBuilder: layerReturnExpectationBuilder,
		layerReturnKindBuilder:        layerReturnKindBuilder,
		layerParamsBuilder:            layerParamsBuilder,
		layerParamBuilder:             layerParamBuilder,
		layerExecutionsBuilder:        layerExecutionsBuilder,
		layerExecutionBuilder:         layerExecutionBuilder,
		layerAssignmentBuilder:        layerAssignmentBuilder,
		layerConditionBuilder:         layerConditionBuilder,
		layerAssignableBuilder:        layerAssignableBuilder,
		layerReduceBuilder:            layerReduceBuilder,
		layerConstantValuesBuilder:    layerConstantValuesBuilder,
		layerConstantValueBuilder:     layerConstantValueBuilder,
		layerQueryBuilder:             layerQueryBuilder,
		layerLayerInputBuilder:        layerLayerInputBuilder,
		layerValueAssignmentsBuilder:  layerValueAssignmentsBuilder,
		layerValueAssignmentBuilder:   layerValueAssignmentBuilder,
		layerValueBuilder:             layerValueBuilder,
	}

	return &out
}

// Execute compiles a resource
func (app *compilerApplication) Execute(input []byte) (symbols.Symbols, error) {
	// init the registry:
	app.pRegistry = app.pRegistry.init()

	// unmarshal the json data:
	ptr := new(map[string]Symbol)
	err := json.Unmarshal(input, ptr)
	if err != nil {
		return nil, err

	}

	// process the symbol list:
	processedList, err := app.processSymbols(*ptr)
	if err != nil {
		return nil, err

	}

	// build the symbols:
	return app.symbolsBuilder.Create().
		WithList(processedList).
		Now()
}

func (app *compilerApplication) processSymbols(input map[string]Symbol) ([]symbols.Symbol, error) {
	processed, remaining, err := app.symbols(input)
	if err != nil {
		return nil, err
	}

	if len(remaining) > 0 {
		list, err := app.processSymbols(remaining)
		if err != nil {
			return nil, err
		}

		processed = append(processed, list...)
	}

	return processed, nil
}

func (app *compilerApplication) symbols(input map[string]Symbol) ([]symbols.Symbol, map[string]Symbol, error) {
	remaining := map[string]Symbol{}
	list := []symbols.Symbol{}
	for name, oneSymbol := range input {
		ins, err := app.symbol(name, oneSymbol)
		if err != nil {
			remaining[name] = oneSymbol
			continue
		}

		err = app.pRegistry.register(name, ins)
		if err != nil {
			return nil, nil, err
		}

		list = append(list, ins)
	}

	if len(list) <= 0 && len(remaining) > 0 {
		namesList := []string{}
		for oneName := range remaining {
			namesList = append(namesList, oneName)
		}

		str := fmt.Sprintf("these symbols (names: %s) could not be processed properly", strings.Join(namesList, ", "))
		return nil, nil, errors.New(str)
	}

	return list, remaining, nil
}

func (app *compilerApplication) symbol(name string, symbol Symbol) (symbols.Symbol, error) {
	builder := app.symbolBuilder.Create()
	if symbol.String != "" {
		builder.WithBytes([]byte(symbol.String))
	}

	if len(symbol.Bytes) > 0 {
		builder.WithBytes(symbol.Bytes)
	}

	if symbol.Layer != nil {
		ins, err := app.layer(*symbol.Layer)
		if err != nil {
			return nil, err
		}

		builder.WithLayer(ins)
	}

	if symbol.Link != nil {
		link, err := app.link(*symbol.Link)
		if err != nil {
			return nil, err
		}

		builder.WithLink(link)
	}

	return builder.Now()
}

func (app *compilerApplication) link(input Link) (links.Link, error) {
	origins, err := app.linkOrigins(input.Origins)
	if err != nil {
		return nil, err
	}

	execution, err := app.linkExecution(input.Execution)
	if err != nil {
		return nil, err
	}

	preparations, err := app.linkPreparations(input.Preparations)
	if err != nil {
		return nil, err
	}

	return app.linkBuilder.Create().
		WithOrigins(origins).
		WithExecution(execution).
		WithPreparations(preparations).
		Now()
}

func (app *compilerApplication) linkOrigins(input []LinkOrigin) (origins.Origins, error) {
	list := []origins.Origin{}
	for _, oneInput := range input {
		ins, err := app.linkOrigin(oneInput)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.linkOriginsBuilder.Create().
		WithList(list).
		Now()
}

func (app *compilerApplication) linkOrigin(input LinkOrigin) (origins.Origin, error) {
	pointer, err := app.pointer(input.Symbol)
	if err != nil {
		return nil, err
	}

	builder := app.linkOriginBuilder.Create().
		WithSymbol(pointer)

	if input.Direction != nil {
		direction, err := app.linkOriginDirection(*input.Direction)
		if err != nil {
			return nil, err
		}

		builder.WithDirection(direction)
	}

	return builder.Now()
}

func (app *compilerApplication) pointer(input Pointer) (pointers.Pointer, error) {
	symbol, err := app.pointerSymbol(input.Symbol)
	if err != nil {
		return nil, err
	}

	return app.pointerBuilder.Create().
		WithPath(input.Path).
		WithSymbol(symbol).
		Now()
}

func (app *compilerApplication) pointerSymbol(input PointerSymbol) (pointer_symbols.Symbol, error) {
	builder := app.pointerSymbolBuilder.Create().
		WithName(input.Name)

	if input.Kind == "bytes" {
		builder.WithKind(pointer_symbols.KindBytes)
	}

	if input.Kind == "layer" {
		builder.WithKind(pointer_symbols.KindLayer)
	}

	if input.Kind == "link" {
		builder.WithKind(pointer_symbols.KindLink)
	}

	return builder.Now()
}

func (app *compilerApplication) linkOriginDirection(input LinkOriginDirection) (origins.Direction, error) {
	builder := app.linkOriginDirectionBuilder.Create()
	if input.Next != nil {
		next, err := app.linkOrigin(*input.Next)
		if err != nil {
			return nil, err
		}

		builder.WithNext(next)
	}

	if input.Previous != nil {
		previous, err := app.linkOrigin(*input.Previous)
		if err != nil {
			return nil, err
		}

		builder.WithPrevious(previous)
	}

	return builder.Now()
}

func (app *compilerApplication) linkExecution(input LinkExecution) (executions.Execution, error) {
	layer, err := app.layerInput(input.Layer)
	if err != nil {
		return nil, err
	}

	builder := app.linkExecutionBuilder.Create().
		WithLayer(layer)

	if len(input.Values) > 0 {
		values, err := app.layerValueAssignments(input.Values)
		if err != nil {
			return nil, err
		}

		builder.WithValues(values)
	}

	return builder.Now()
}

func (app *compilerApplication) linkPreparations(input []LinkPreparation) (preparations.Preparations, error) {
	list := []preparations.Preparation{}
	for _, oneInput := range input {
		ins, err := app.linkPreparation(oneInput)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.linkPreparationsBuilder.Create().
		WithList(list).
		Now()
}

func (app *compilerApplication) linkPreparation(input LinkPreparation) (preparations.Preparation, error) {
	builder := app.linkPreparationBuilder.Create()
	if input.IsStop {
		builder.IsStop()
	}

	if input.Load != nil {
		pointer, err := app.pointer(*input.Load)
		if err != nil {
			return nil, err
		}

		builder.WithLoad(pointer)
	}

	if input.Exists != nil {
		pointer, err := app.pointer(*input.Exists)
		if err != nil {
			return nil, err
		}

		builder.WithExists(pointer)
	}

	if input.Condition != nil {
		condition, err := app.linkPreparationCondition(*input.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *compilerApplication) linkPreparationCondition(input LinkPreparationCondition) (preparations.Condition, error) {
	preparations, err := app.linkPreparations(input.Preparations)
	if err != nil {
		return nil, err
	}

	return app.linkConditionBuilder.Create().
		WithVariable(input.Variable).
		WithPreparations(preparations).
		Now()
}

func (app *compilerApplication) layer(input Layer) (layers.Layer, error) {
	executions, err := app.layerExecutions(input.Executions)
	if err != nil {
		return nil, err
	}

	ret, err := app.layerReturnExpectation(input.Return)
	if err != nil {
		return nil, err
	}

	builder := app.layerBuilder.Create().
		WithExecutions(executions).
		WithReturn(ret).
		WithInput(input.Input)

	if len(input.Parameters) > 0 {
		parameters, err := app.layerParameters(input.Parameters)
		if err != nil {
			return nil, err
		}

		builder.WithParams(parameters)
	}

	return builder.Now()
}

func (app *compilerApplication) layerExecutions(input []LayerExecution) (layers.Executions, error) {
	list := []layers.Execution{}
	for _, oneInput := range input {
		ins, err := app.layerExecution(oneInput)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.layerExecutionsBuilder.Create().
		WithList(list).
		Now()
}

func (app *compilerApplication) layerExecution(input LayerExecution) (layers.Execution, error) {
	builder := app.layerExecutionBuilder.Create()
	if input.IsStop {
		builder.IsStop()
	}

	if input.Assignment != nil {
		assignment, err := app.layerExecutionAssignment(*input.Assignment)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(assignment)
	}

	if input.Condition != nil {
		condition, err := app.layerExecutionCondition(*input.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *compilerApplication) layerExecutionAssignment(input LayerExecutionAssignment) (layers.Assignment, error) {
	assignable, err := app.layerExecutionAssignable(input.Assignable)
	if err != nil {
		return nil, err
	}

	return app.layerAssignmentBuilder.Create().
		WithName(input.Name).
		WithAssignable(assignable).
		Now()
}

func (app *compilerApplication) layerExecutionAssignable(input LayerExecutionAssignable) (layers.Assignable, error) {
	builder := app.layerAssignableBuilder.Create()
	if input.Query != nil {
		query, err := app.layerExecutionQuery(*input.Query)
		if err != nil {
			return nil, err
		}

		builder.WithQuery(query)
	}

	if input.Reduce != nil {
		reduce, err := app.layerExecutionReduce(*input.Reduce)
		if err != nil {
			return nil, err
		}

		builder.WithReduce(reduce)
	}

	if len(input.Compare) > 0 {
		compare, err := app.layerConstantValues(input.Compare)
		if err != nil {
			return nil, err
		}

		builder.WithCompare(compare)
	}

	if input.Length != nil {
		length, err := app.layerConstantValue(*input.Length)
		if err != nil {
			return nil, err
		}

		builder.WithLength(length)
	}

	if len(input.Join) > 0 {
		join, err := app.layerConstantValues(input.Join)
		if err != nil {
			return nil, err
		}

		builder.WithJoin(join)
	}

	if input.Value != nil {
		value, err := app.layerConstantValue(*input.Value)
		if err != nil {
			return nil, err
		}

		builder.WithValue(value)
	}

	if input.Administrator != nil {

	}

	return builder.Now()
}

func (app *compilerApplication) administrator(input administrators.Administrator) error {
	return nil
}

func (app *compilerApplication) layerExecutionQuery(input LayerExecutionQuery) (layers.Query, error) {
	inputIns, err := app.layerConstantValue(input.Input)
	if err != nil {
		return nil, err
	}

	layer, err := app.layerInput(input.Layer)
	if err != nil {
		return nil, err
	}

	builder := app.layerQueryBuilder.Create().
		WithInput(inputIns).
		WithLayer(layer)

	if len(input.Values) > 0 {
		values, err := app.layerValueAssignments(input.Values)
		if err != nil {
			return nil, err
		}

		builder.WithValues(values)
	}

	return builder.Now()
}

func (app *compilerApplication) layerInput(input LayerInput) (layers.LayerInput, error) {
	builder := app.layerLayerInputBuilder.Create()
	if input.Variable != "" {
		builder.WithVariable(input.Variable)
	}

	if input.Layer != "" {

	}

	return builder.Now()
}

func (app *compilerApplication) layerValueAssignments(input []LayerValueAssignment) (layers.ValueAssignments, error) {
	list := []layers.ValueAssignment{}
	for _, oneInput := range input {
		ins, err := app.layerValueAssignment(oneInput)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.layerValueAssignmentsBuilder.Create().
		WithList(list).
		Now()
}

func (app *compilerApplication) layerValueAssignment(input LayerValueAssignment) (layers.ValueAssignment, error) {
	value, err := app.layerValue(input.Value)
	if err != nil {
		return nil, err
	}

	return app.layerValueAssignmentBuilder.Create().
		WithName(input.Name).
		WithValue(value).
		Now()
}

func (app *compilerApplication) layerValue(input LayerValue) (layers.Value, error) {
	builder := app.layerValueBuilder.Create()
	if input.Variable != "" {
		builder.WithVariable(input.Variable)
	}

	if len(input.Bytes) > 0 {
		symbol, err := app.pRegistry.retrieve(input.Bytes)
		if err != nil {
			return nil, err
		}

		if !symbol.IsBytes() {
			str := fmt.Sprintf("the Symbol (name: %s) was expected to contain []byte", input.Layer)
			return nil, errors.New(str)
		}

		bytes := symbol.Bytes()
		builder.WithConstant(bytes)
	}

	if input.Layer != "" {
		symbol, err := app.pRegistry.retrieve(input.Layer)
		if err != nil {
			return nil, err
		}

		if !symbol.IsLayer() {
			str := fmt.Sprintf("the Symbol (name: %s) was expected to contain a Layer instance", input.Layer)
			return nil, errors.New(str)
		}

		layer := symbol.Layer()
		builder.WithLayer(layer)
	}

	return builder.Now()
}

func (app *compilerApplication) layerExecutionReduce(input LayerExecutionReduce) (reduces.Reduce, error) {
	return app.layerReduceBuilder.Create().
		WithVariable(input.Variable).
		WithLength(input.Length).
		Now()
}

func (app *compilerApplication) layerConstantValues(input []LayerConstantValue) (constantvalues.ConstantValues, error) {
	list := []constantvalues.ConstantValue{}
	for _, oneInput := range input {
		ins, err := app.layerConstantValue(oneInput)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.layerConstantValuesBuilder.Create().
		WithList(list).
		Now()
}

func (app *compilerApplication) layerConstantValue(input LayerConstantValue) (constantvalues.ConstantValue, error) {
	builder := app.layerConstantValueBuilder.Create()
	if input.Variable != "" {
		builder.WithVariable(input.Variable)
	}

	if len(input.Bytes) > 0 {
		symbol, err := app.pRegistry.retrieve(input.Bytes)
		if err != nil {
			return nil, err
		}

		if !symbol.IsBytes() {
			str := fmt.Sprintf("the Symbol (name: %s) was expected to contain []byte", input.Bytes)
			return nil, errors.New(str)
		}

		bytes := symbol.Bytes()
		builder.WithConstant(bytes)
	}

	return builder.Now()
}

func (app *compilerApplication) layerExecutionCondition(input LayerExecutionCondition) (layers.Condition, error) {
	executions, err := app.layerExecutions(input.Executions)
	if err != nil {
		return nil, err
	}

	return app.layerConditionBuilder.Create().
		WithVariable(input.Variable).
		WithExecutions(executions).
		Now()
}

func (app *compilerApplication) layerReturnExpectation(input LayerReturn) (return_expectations.Expectation, error) {
	kind, err := app.layerReturnKind(input.Kind)
	if err != nil {
		return nil, err
	}

	return app.layerReturnExpectationBuilder.Create().
		WithVariable(input.Variable).
		WithKind(kind).
		Now()
}

func (app *compilerApplication) layerReturnKind(input LayerReturnKind) (kinds.Kind, error) {
	builder := app.layerReturnKindBuilder.Create()
	if input.IsContinue {
		builder.IsContinue()
	}

	if input.IsPrompt {
		builder.IsPrompt()
	}

	if len(input.Execute) > 0 {
		builder.WithExecute(input.Execute)
	}

	return builder.Now()
}

func (app *compilerApplication) layerParameters(input map[string]string) (parameters.Parameters, error) {
	list := []parameters.Parameter{}
	for name, kindString := range input {
		kind := uint8(0)
		if kindString == "bytes" {
			kind = pointer_symbols.KindBytes
		}

		if kindString == "layer" {
			kind = pointer_symbols.KindLayer
		}

		ins, err := app.layerParameter(name, kind)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.layerParamsBuilder.Create().
		WithList(list).
		Now()
}

func (app *compilerApplication) layerParameter(name string, kind uint8) (parameters.Parameter, error) {
	return app.layerParamBuilder.Create().
		WithName(name).
		WithKind(kind).
		Now()
}

func (app *compilerApplication) value(input struct_values.Value) []byte {
	if input.String != "" {
		return []byte(input.String)
	}

	return input.Bytes
}
