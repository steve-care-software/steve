package administrators

import (
	"github.com/steve-care-software/steve/applications/administrators/administrators/authenticates"
	"github.com/steve-care-software/steve/applications/administrators/administrators/instances"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators"
)

type application struct {
	authApp         authenticates.Application
	instanceApp     instances.Application
	exectionBuilder executions.Builder
}

func createApplication(
	authApp authenticates.Application,
	instanceApp instances.Application,
	exectionBuilder executions.Builder,
) Application {
	out := application{
		authApp:         authApp,
		instanceApp:     instanceApp,
		exectionBuilder: exectionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(instance inputs.Administrator) (executions.Administrator, error) {
	builder := app.exectionBuilder.Create()
	if instance.IsAuthenticate() {
		auth := instance.Authenticate()
		exec, err := app.authApp.Execute(auth)
		if err != nil {
			return nil, err
		}

		builder.WithAuthenticate(exec)
	}

	if instance.IsInstance() {
		ins := instance.Instance()
		exec, err := app.instanceApp.Execute(ins)
		if err != nil {
			return nil, err
		}

		builder.WithInstance(exec)
	}

	return builder.Now()
}
