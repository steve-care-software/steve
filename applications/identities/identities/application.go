package identities

import (
	"github.com/steve-care-software/steve/applications/identities/identities/authenticates"
	"github.com/steve-care-software/steve/applications/identities/identities/instances"
	"github.com/steve-care-software/steve/applications/identities/identities/lists"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities"
	"github.com/steve-care-software/steve/domain/stacks"
)

type application struct {
	authenticateApp  authenticates.Application
	instanceApp      instances.Application
	listApp          lists.Application
	executionBuilder executions.Builder
}

func createApplication(
	authenticateApp authenticates.Application,
	instanceApp instances.Application,
	listApp lists.Application,
	executionBuilder executions.Builder,
) Application {
	out := application{
		authenticateApp:  authenticateApp,
		instanceApp:      instanceApp,
		listApp:          listApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(identity inputs.Identity, stack stacks.Stack) (executions.Identity, error) {
	builder := app.executionBuilder.Create()
	if identity.IsAuthenticate() {
		authenticate := identity.Authenticate()
		exec, err := app.authenticateApp.Execute(authenticate)
		if err != nil {
			return nil, err
		}

		builder.WithAuthenticate(exec)
	}

	if identity.IsList() {
		list := identity.List()
		exec, err := app.listApp.Execute(list, stack)
		if err != nil {
			return nil, err
		}

		builder.WithList(exec)
	}

	if identity.IsInstance() {
		instance := identity.Instance()
		exec, err := app.instanceApp.Execute(instance, stack)
		if err != nil {
			return nil, err
		}

		builder.WithInstance(exec)
	}

	return builder.Now()
}
