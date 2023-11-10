package administrators

import (
	"github.com/steve-care-software/steve/applications/commands/shares/administrators/creates"
	execution_administrators "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/administrators"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/administrators"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/shares/administrators"
)

type application struct {
	createApp        creates.Application
	executionBuilder execution_administrators.Builder
}

func createApplication(
	createApp creates.Application,
	executionBuilder execution_administrators.Builder,
) Application {
	out := application{
		createApp:        createApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(administrator inputs.Administrator) (executions.Administrator, error) {
	builder := app.executionBuilder.Create()
	if administrator.IsCreate() {
		create := administrator.Create()
		exec, err := app.createApp.Execute(create)
		if err != nil {
			return nil, err
		}

		builder.WithCreate(exec)
	}

	return builder.Now()
}
