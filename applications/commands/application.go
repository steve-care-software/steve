package commands

import (
	"github.com/steve-care-software/steve/applications/commands/administrators"
	"github.com/steve-care-software/steve/applications/commands/visitors"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs"
)

type application struct {
	visitorApp       visitors.Application
	adminApp         administrators.Application
	inputAdapter     inputs.Adapter
	executionBuilder executions.Builder
}

func createApplication(
	visitorApp visitors.Application,
	adminApp administrators.Application,
	inputAdapter inputs.Adapter,
	executionBuilder executions.Builder,
) Application {
	out := application{
		visitorApp:       visitorApp,
		adminApp:         adminApp,
		inputAdapter:     inputAdapter,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(input []byte) (executions.Execution, error) {
	inputIns, err := app.inputAdapter.ToInput(input)
	if err != nil {
		return nil, err
	}

	builder := app.executionBuilder.Create()
	if inputIns.IsAdministrator() {
		admin := inputIns.Administrator()
		retExec, err := app.adminApp.Execute(admin, nil)
		if err != nil {
			return nil, err
		}

		builder.WithAdministrator(retExec)
	}

	if inputIns.IsVisitor() {
		visitor := inputIns.Visitor()
		retExec, err := app.visitorApp.Execute(visitor)
		if err != nil {
			return nil, err
		}

		builder.WithVisitor(retExec)
	}

	return builder.Now()
}
