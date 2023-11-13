package commands

import (
	"github.com/steve-care-software/steve/applications/commands/administrators"
	"github.com/steve-care-software/steve/applications/commands/visitors"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs"
)

type application struct {
	visitorApp       visitors.Application
	adminApp         administrators.Application
	executionBuilder executions.Builder
}

func createApplication(
	visitorApp visitors.Application,
	adminApp administrators.Application,
	executionBuilder executions.Builder,
) Application {
	out := application{
		visitorApp:       visitorApp,
		adminApp:         adminApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(inputIns inputs.Input, frame frames.Frame) (executions.Execution, error) {
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
