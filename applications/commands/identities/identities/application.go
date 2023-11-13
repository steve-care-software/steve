package identities

import (
	"github.com/steve-care-software/steve/applications/commands/identities/identities/authenticates"
	"github.com/steve-care-software/steve/applications/commands/identities/identities/instances"
	"github.com/steve-care-software/steve/applications/commands/identities/identities/lists"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/identities"
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
func (app *application) Execute(identity inputs.Identity, frame frames.Frame) (executions.Identity, error) {
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
		exec, err := app.listApp.Execute(list, frame)
		if err != nil {
			return nil, err
		}

		builder.WithList(exec)
	}

	if identity.IsInstance() {
		instance := identity.Instance()
		exec, err := app.instanceApp.Execute(instance, frame)
		if err != nil {
			return nil, err
		}

		builder.WithInstance(exec)
	}

	return builder.Now()
}
