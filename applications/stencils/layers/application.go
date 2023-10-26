package layers

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
	"github.com/steve-care-software/steve/domain/stencils/messages"
	"github.com/steve-care-software/steve/domain/stencils/queries"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/executions"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/executions/assignments"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/inits"

	pointer_symbols "github.com/steve-care-software/steve/domain/stencils/pointers/symbols"
	result_layers "github.com/steve-care-software/steve/domain/stencils/results/executions"
	init_inputs "github.com/steve-care-software/steve/domain/stencils/results/executions/inits/inputs"
	init_values "github.com/steve-care-software/steve/domain/stencils/results/executions/inits/values"
)

type application struct {
	computer           *computer
	resultLayerBuilder result_layers.ExecutionBuilder
	initBuilder        inits.Builder
	initInputBuilder   init_inputs.Builder
	initValuesBuilder  init_values.Builder
	initValueBuilder   init_values.ValueBuilder
	executionsBuilder  executions.Builder
	executionBuilder   executions.ExecutionBuilder
	assignmentBulder   assignments.Builder
	queryBuilder       queries.Builder
	messageBuilder     messages.Builder
}

func createApplication(
	computer *computer,
) Application {
	out := application{
		computer: computer,
	}

	return &out
}

// Execute executes a layer
func (app *application) Execute(query queries.Query) (result_layers.Execution, error) {
	layerIns := query.Layer()
	init, err := app.init(layerIns, query)
	if err != nil {
		return nil, err
	}

	pContext, err := app.computer.init(init)
	if err != nil {
		return nil, err
	}

	retExecutions, err := app.executions(*pContext, layerIns.Executions())
	if err != nil {
		return nil, err
	}

	ret := layerIns.Return()
	return app.resultLayerBuilder.Create().
		WithInit(init).
		WithExecutions(retExecutions).
		WithReturn(ret).
		Now()
}

func (app *application) executions(context uint, layerExecutions layers.Executions) (executions.Executions, error) {
	list := layerExecutions.List()
	output := []executions.Execution{}
	for idx, oneExecution := range list {
		execution, err := app.execution(context, uint(idx), oneExecution)
		if err != nil {
			return nil, err
		}

		output = append(output, execution)
	}

	return app.executionsBuilder.Create().
		WithList(output).
		Now()
}

func (app *application) execution(context uint, idx uint, execution layers.Execution) (executions.Execution, error) {
	builder := app.executionBuilder.Create()
	if execution.IsStop() {
		builder.IsStop()
	}

	if execution.IsAssignment() {
		assignment := execution.Assignment()
		retAssignment, err := app.assignment(context, assignment)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(retAssignment)
	}

	if execution.IsCondition() {
		condition := execution.Condition()
		retExecutions, err := app.condition(context, condition)
		if err != nil {
			return nil, err
		}

		if retExecutions == nil {
			return nil, nil
		}

		builder.WithExecutions(retExecutions)
	}

	return builder.Now()
}

func (app *application) condition(context uint, condition layers.Condition) (executions.Executions, error) {
	variable := condition.Variable()
	values, err := app.computer.retrieve(context, variable)
	if err != nil {
		return nil, err
	}

	if len(values) != 1 {
		str := fmt.Sprintf("the variable (name: %s) is used in a condition an was therefore expected to contain a bool value", variable)
		return nil, errors.New(str)
	}

	if values[0] == 0 {
		executions := condition.Executions()
		return app.executions(context, executions)
	}

	return nil, nil
}

func (app *application) assignment(context uint, assignment layers.Assignment) (assignments.Assignment, error) {
	name := assignment.Name()
	assignable := assignment.Assignable()
	values, err := app.assignable(context, assignable)
	if err != nil {
		return nil, err
	}

	err = app.computer.assign(context, name, values)
	if err != nil {
		return nil, err
	}

	return app.assignmentBulder.Create().
		WithName(name).
		WithValue(values).
		Now()
}

func (app *application) assignable(context uint, assignable layers.Assignable) ([]byte, error) {
	if assignable.IsQuery() {
		query := assignable.Query()
		criteria, err := app.query(context, query)
		if err != nil {
			return nil, err
		}

		result, err := app.Execute(criteria)
		if err != nil {
			return nil, err
		}

		return result.Bytes(), nil
	}

	if assignable.IsReduce() {
		reduce := assignable.Reduce()
		return app.reduce(context, reduce)
	}

	if assignable.IsCompare() {
		cmp := assignable.Compare()
		return app.compare(context, cmp)
	}

	if assignable.IsLength() {
		length := assignable.Length()
		return app.length(context, length)
	}

	if assignable.IsJoin() {
		join := assignable.Join()
		return app.join(context, join)
	}

	value := assignable.Value()
	return app.value(context, value)
}

func (app *application) query(context uint, query layers.Query) (queries.Query, error) {
	inputValue := query.Input()
	input, err := app.value(context, inputValue)
	if err != nil {
		return nil, err
	}

	layerInput := query.Layer()
	layer, err := app.layerInput(context, layerInput)
	if err != nil {
		return nil, err
	}

	msg, err := app.messageBuilder.Create().
		WithBytes(input).
		Now()

	if err != nil {
		return nil, err
	}

	builder := app.queryBuilder.Create().
		WithMessage(msg).
		WithLayer(layer)

	if query.HasValues() {
		valueAssignments := query.Values()
		params, err := app.valueAssignments(context, valueAssignments)
		if err != nil {
			return nil, err
		}

		builder.WithParams(params)
	}

	return builder.Now()
}

func (app *application) layerInput(context uint, input layers.LayerInput) (layers.Layer, error) {
	if input.IsLayer() {
		return input.Layer(), nil
	}

	name := input.Variable()
	return app.computer.retrieveLayer(context, name)
}

func (app *application) valueAssignments(context uint, valueAssignments layers.ValueAssignments) (symbols.Symbols, error) {
	return nil, nil
}

func (app *application) valueAssignment(context uint, valueAssignment layers.ValueAssignment) (symbols.Symbol, error) {
	return nil, nil
}

/*

Name() string
	Value() Value

IsVariable() bool
	Variable() string
	IsLayer() bool
	Layer() Layer

*/

func (app *application) reduce(context uint, reduce reduces.Reduce) ([]byte, error) {
	variable := reduce.Variable()
	length := reduce.Length()
	return app.computer.reduce(context, variable, length)
}

func (app *application) compare(context uint, constants layers.ConstantValues) ([]byte, error) {
	first := []byte{}
	list := constants.List()
	for _, oneConstant := range list {
		values, err := app.value(context, oneConstant)
		if err != nil {
			return nil, err
		}

		if len(first) <= 0 {
			first = values
			continue
		}

		if bytes.Equal(first, values) {
			continue
		}

		return []byte{1}, nil
	}

	return []byte{0}, nil
}

func (app *application) length(context uint, constant layers.ConstantValue) ([]byte, error) {
	bytes, err := app.value(context, constant)
	if err != nil {
		return nil, err
	}

	length := len(bytes)
	return []byte(strconv.Itoa(length)), nil
}

func (app *application) join(context uint, constants layers.ConstantValues) ([]byte, error) {
	output := []byte{}
	list := constants.List()
	for _, oneConstant := range list {
		bytes, err := app.value(context, oneConstant)
		if err != nil {
			return nil, err
		}

		output = append(output, bytes...)
	}

	return output, nil
}

func (app *application) value(context uint, constant layers.ConstantValue) ([]byte, error) {
	if constant.IsConstant() {
		return constant.Constant(), nil
	}

	name := constant.Variable()
	return app.computer.retrieve(context, name)
}

func (app *application) init(layerIns layers.Layer, query queries.Query) (inits.Init, error) {
	variable := layerIns.Input()
	inputBytes := query.Message().Bytes()
	inputIns, err := app.initInputBuilder.Create().
		WithVariable(variable).
		WithBytes(inputBytes).
		Now()

	if err != nil {
		return nil, err
	}

	ret := layerIns.Return()
	initBuilder := app.initBuilder.Create().
		WithInput(inputIns).
		WithReturn(ret)

	if query.HasParams() {
		paramSymbols := query.Params()
		layerParams := layerIns.Params()
		inputValues, err := app.inputValues(paramSymbols, layerParams)
		if err != nil {
			return nil, err
		}

		initBuilder.WithValues(inputValues)
	}

	return initBuilder.Now()
}

func (app *application) inputValues(symbolsIns symbols.Symbols, params parameters.Parameters) (init_values.Values, error) {
	list := []init_values.Value{}
	paramsList := params.List()
	for _, oneParameter := range paramsList {
		expectedKind := oneParameter.Kind()
		name := oneParameter.Name()
		symbol, err := symbolsIns.Fetch(name)
		if err != nil {
			return nil, err
		}

		valueBuilder := app.initValueBuilder.Create().
			WithVariable(name)

		if expectedKind == pointer_symbols.KindBytes {
			if !symbol.IsBytes() {
				str := fmt.Sprintf("the parameter (name: %s) was expected to contain bytes", name)
				return nil, errors.New(str)
			}

			bytes := symbol.Bytes()
			valueBuilder.WithBytes(bytes)
		}

		if expectedKind == pointer_symbols.KindLayer {
			if !symbol.IsLayer() {
				str := fmt.Sprintf("the parameter (name: %s) was expected to contain a layer", name)
				return nil, errors.New(str)
			}

			layer := symbol.Layer()
			valueBuilder.WithLayer(layer)
		}

		value, err := valueBuilder.Now()
		if err != nil {
			return nil, err
		}

		list = append(list, value)
	}

	return app.initValuesBuilder.Create().
		WithList(list).
		Now()
}
