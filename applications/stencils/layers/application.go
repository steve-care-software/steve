package layers

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	account_applications "github.com/steve-care-software/steve/applications/accounts"
	applications_administrator "github.com/steve-care-software/steve/applications/accounts/administrators"
	"github.com/steve-care-software/steve/domain/credentials"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
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
	accountApp                      account_applications.Application
	accountAdminSaveCriteriaBuilder applications_administrator.SaveCriteriaBuilder
	resultLayerBuilder              result_layers.ExecutionBuilder
	initBuilder                     inits.Builder
	initInputBuilder                init_inputs.Builder
	initValuesBuilder               init_values.Builder
	initValueBuilder                init_values.ValueBuilder
	executionsBuilder               executions.Builder
	executionBuilder                executions.ExecutionBuilder
	assignmentBulder                assignments.Builder
	queryBuilder                    queries.Builder
	messageBuilder                  messages.Builder
	credentialsBuilder              credentials.Builder
	valueBuilder                    ValueBuilder
	valueAccountBuilder             AccountBuilder
	values                          map[string]Value
}

func createApplication() Application {
	out := application{
		values: map[string]Value{},
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(query queries.Query) (result_layers.Execution, error) {
	layerIns := query.Layer()
	init, err := app.init(layerIns, query)
	if err != nil {
		return nil, err
	}

	err = app.reset(init)
	if err != nil {
		return nil, err
	}

	retExecutions, err := app.executions(layerIns.Executions())
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

func (app *application) reset(init inits.Init) error {
	app.values = map[string]Value{}
	input := init.Input()
	inputValue, err := app.valueBuilder.Create().WithBytes(input.Bytes()).Now()
	if err != nil {
		return err
	}

	inputVariable := input.Variable()
	app.assign(inputVariable, inputValue)
	if init.HasValues() {
		values := init.Values().List()
		for _, oneValue := range values {
			content := oneValue.Content()
			valueBuilder := app.valueBuilder.Create()
			if content.IsBytes() {
				bytes := content.Bytes()
				valueBuilder.WithBytes(bytes)
			}

			if content.IsLayer() {
				layer := content.Layer()
				valueBuilder.WithLayer(layer)
			}

			value, err := valueBuilder.Now()
			if err != nil {
				return err
			}

			variable := oneValue.Variable()
			app.assign(variable, value)
		}
	}

	return nil
}

func (app *application) retrieve(name string) (Value, error) {
	if ins, ok := app.values[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the variable (name: %s) cannot be retrieved because it has never been assigned", name)
	return nil, errors.New(str)
}

func (app *application) assign(name string, value Value) {
	app.values[name] = value
}

func (app *application) executions(layerExecutions layers.Executions) (executions.Executions, error) {
	list := layerExecutions.List()
	output := []executions.Execution{}
	for idx, oneExecution := range list {
		execution, err := app.execution(uint(idx), oneExecution)
		if err != nil {
			return nil, err
		}

		output = append(output, execution)
	}

	return app.executionsBuilder.Create().
		WithList(output).
		Now()
}

func (app *application) execution(idx uint, execution layers.Execution) (executions.Execution, error) {
	builder := app.executionBuilder.Create()
	if execution.IsStop() {
		builder.IsStop()
	}

	if execution.IsAssignment() {
		assignment := execution.Assignment()
		retAssignment, err := app.assignment(assignment)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(retAssignment)
	}

	if execution.IsCondition() {
		condition := execution.Condition()
		retExecutions, err := app.condition(condition)
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

func (app *application) condition(condition layers.Condition) (executions.Executions, error) {
	variable := condition.Variable()
	value, err := app.retrieve(variable)
	if err != nil {
		return nil, err
	}

	if !value.IsBytes() {
		str := fmt.Sprintf("the variable (name: %s) is used in a condition an was therefore expected to contain a bool value (bytes of length 1)", variable)
		return nil, errors.New(str)
	}

	bytes := value.Bytes()
	if len(bytes) != 1 {
		str := fmt.Sprintf("the variable (name: %s) is used in a condition an was therefore expected to contain a bool value (bytes of length 1), %d bytes provided", variable, len(bytes))
		return nil, errors.New(str)
	}

	if bytes[0] == 0 {
		executions := condition.Executions()
		return app.executions(executions)
	}

	return nil, nil
}

func (app *application) assignment(assignment layers.Assignment) (assignments.Assignment, error) {
	name := assignment.Name()
	assignable := assignment.Assignable()
	value, err := app.assignable(assignable)
	if err != nil {
		return nil, err
	}

	app.assign(name, value)
	builder := app.assignmentBulder.Create().WithName(name)
	if value.IsBytes() {
		bytes := value.Bytes()
		builder.WithValue(bytes)
	}

	if !value.IsBytes() {
		builder.IsInternal()
	}

	return builder.Now()
}

func (app *application) assignable(assignable layers.Assignable) (Value, error) {
	valueBuilder := app.valueBuilder.Create()
	if assignable.IsQuery() {
		query := assignable.Query()
		criteria, err := app.query(query)
		if err != nil {
			return nil, err
		}

		result, err := app.Execute(criteria)
		if err != nil {
			return nil, err
		}

		valueBuilder.WithResult(result)
	}

	if assignable.IsReduce() {
		reduce := assignable.Reduce()
		bytes, err := app.reduce(reduce)
		if err != nil {
			return nil, err
		}

		valueBuilder.WithBytes(bytes)
	}

	if assignable.IsCompare() {
		cmp := assignable.Compare()
		bytes, err := app.compare(cmp)
		if err != nil {
			return nil, err
		}

		valueBuilder.WithBytes(bytes)
	}

	if assignable.IsLength() {
		length := assignable.Length()
		bytes, err := app.length(length)
		if err != nil {
			return nil, err
		}

		valueBuilder.WithBytes(bytes)
	}

	if assignable.IsJoin() {
		join := assignable.Join()
		bytes, err := app.join(join)
		if err != nil {
			return nil, err
		}

		valueBuilder.WithBytes(bytes)
	}

	if assignable.IsValue() {
		value := assignable.Value()
		bytes, err := app.value(value)
		if err != nil {
			return nil, err
		}

		valueBuilder.WithBytes(bytes)
	}

	if assignable.IsAccount() {
		account := assignable.Account()
		retAccount, err := app.account(account)
		if err != nil {
			return nil, err
		}

		valueBuilder.WithAccount(retAccount)
	}

	return valueBuilder.Now()
}

func (app *application) account(account accounts.Account) (Account, error) {
	builder := app.valueAccountBuilder.Create()
	if account.IsAdministrator() {
		administrator := account.Administrator()
		variable := administrator.Application()
		administratorApp, err := app.administratorApplicationByConstantValue(variable)
		if err != nil {
			return nil, err
		}

		content := administrator.Content()
		if content.IsApplication() {
			adminAppInstruction := content.Application()
			if adminAppInstruction.IsRetrieve() {
				retAdminIns, err := administratorApp.Retrieve()
				if err != nil {
					return nil, err
				}

				builder.WithAdministrator(retAdminIns)
			}

			if adminAppInstruction.IsSave() {
				save := adminAppInstruction.Save()
				instance := save.Instance()
				password, err := app.value(save.Password())
				if err != nil {
					return nil, err
				}

				criteriaBuilder := app.accountAdminSaveCriteriaBuilder.Create().
					WithAdministrator(instance).
					WithPassword(password)

				if save.HasNewPassword() {
					newPassword, err := app.value(save.NewPassword())
					if err != nil {
						return nil, err
					}

					criteriaBuilder.WithNewPassword(newPassword)
				}

				ins, err := criteriaBuilder.Now()
				if err != nil {
					return nil, err
				}

				err = administratorApp.Save(ins)
				if err != nil {
					return nil, err
				}

				panic(errors.New("make save a non-assignable execution"))
			}
		}

		if content.IsInstance() {
			panic(errors.New("finish instance"))
		}
	}

	if account.IsAuthenticate() {
		accCredentials := account.Authenticate()
		username, err := app.value(accCredentials.Username())
		if err != nil {
			return nil, err
		}

		password, err := app.value(accCredentials.Password())
		if err != nil {
			return nil, err
		}

		credentials, err := app.credentialsBuilder.Create().
			WithUsername(string(username)).
			WithPassword(password).
			Now()

		if err != nil {
			return nil, err
		}

		adminApp, err := app.accountApp.Authenticate(credentials)
		if err != nil {
			return nil, err
		}

		builder.WithApplication(adminApp)
	}

	return builder.Now()
}

func (app *application) query(query layers.Query) (queries.Query, error) {
	inputValue := query.Input()
	input, err := app.value(inputValue)
	if err != nil {
		return nil, err
	}

	layerInput := query.Layer()
	layer, err := app.layerInput(layerInput)
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
		params, err := app.valueAssignments(valueAssignments)
		if err != nil {
			return nil, err
		}

		builder.WithParams(params)
	}

	return builder.Now()
}

func (app *application) valueAssignments(valueAssignments layers.ValueAssignments) (symbols.Symbols, error) {
	return nil, nil
}

func (app *application) valueAssignment(valueAssignment layers.ValueAssignment) (symbols.Symbol, error) {
	return nil, nil
}

func (app *application) reduce(reduce reduces.Reduce) ([]byte, error) {
	variable := reduce.Variable()
	length := reduce.Length()
	value, err := app.retrieve(variable)
	if err != nil {
		return nil, err
	}

	if !value.IsBytes() {
		str := fmt.Sprintf("the variable (name: %s) was expected to contain bytes when used using the reduce func")
		return nil, errors.New(str)
	}

	bytes := value.Bytes()
	if len(bytes) < int(length) {
		str := fmt.Sprintf("the variable (name: %s) was expected to be reduced to %d bytes, but it only contains %d bytes", variable, len(bytes), length)
		return nil, errors.New(str)
	}

	return bytes[:int(length)], nil
}

func (app *application) compare(constants constantvalues.ConstantValues) ([]byte, error) {
	first := []byte{}
	list := constants.List()
	for _, oneConstant := range list {
		values, err := app.value(oneConstant)
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

func (app *application) length(constant constantvalues.ConstantValue) ([]byte, error) {
	bytes, err := app.value(constant)
	if err != nil {
		return nil, err
	}

	length := len(bytes)
	return []byte(strconv.Itoa(length)), nil
}

func (app *application) join(constants constantvalues.ConstantValues) ([]byte, error) {
	output := []byte{}
	list := constants.List()
	for _, oneConstant := range list {
		bytes, err := app.value(oneConstant)
		if err != nil {
			return nil, err
		}

		output = append(output, bytes...)
	}

	return output, nil
}

func (app *application) layerInput(input layers.LayerInput) (layers.Layer, error) {
	if input.IsLayer() {
		return input.Layer(), nil
	}

	name := input.Variable()
	retValue, err := app.retrieve(name)
	if err != nil {
		return nil, err
	}

	if !retValue.IsLayer() {
		str := fmt.Sprintf("the variable (%s) was expected to contain a Layer instance", name)
		return nil, errors.New(str)
	}

	return retValue.Layer(), nil
}

func (app *application) administratorApplicationByConstantValue(input constantvalues.ConstantValue) (applications_administrator.Application, error) {
	return nil, nil
}

func (app *application) accountByConstantValue(input constantvalues.ConstantValue) (Account, error) {
	return nil, nil
}

func (app *application) value(constant constantvalues.ConstantValue) ([]byte, error) {
	if constant.IsConstant() {
		return constant.Constant(), nil
	}

	name := constant.Variable()
	retValue, err := app.retrieve(name)
	if err != nil {
		return nil, err
	}

	if !retValue.IsBytes() {
		str := fmt.Sprintf("the variable (%s) was expected to contain bytes", name)
		return nil, errors.New(str)
	}

	return retValue.Bytes(), nil
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
