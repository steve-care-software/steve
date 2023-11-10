package administrators

import (
	"github.com/steve-care-software/steve/applications/commands/administrators/administrators"
	"github.com/steve-care-software/steve/applications/commands/administrators/identities"
	"github.com/steve-care-software/steve/applications/commands/shares/dashboards"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators"
)

type application struct {
	adminApp         administrators.Application
	identitiesApp    identities.Application
	dashboardApp     dashboards.Application
	executionBuilder executions.Builder
}

func createApplication(
	adminApp administrators.Application,
	identitiesApp identities.Application,
	dashboardApp dashboards.Application,
	executionBuilder executions.Builder,
) Application {
	out := application{
		adminApp:         adminApp,
		identitiesApp:    identitiesApp,
		dashboardApp:     dashboardApp,
		executionBuilder: executionBuilder,
	}
	return &out
}

// Execute executes an administrator
func (app *application) Execute(administrator inputs.Administrator, frame frames.Frame) (executions.Administrator, error) {
	builder := app.executionBuilder.Create()
	if administrator.IsAdministrator() {
		admin := administrator.Administrator()
		exec, err := app.adminApp.Execute(admin, frame)
		if err != nil {
			return nil, err
		}

		builder.WithAdministrator(exec)
	}

	if administrator.IsIdentities() {
		identities := administrator.Identities()
		exec, err := app.identitiesApp.Execute(identities, frame)
		if err != nil {
			return nil, err
		}

		builder.WithIdentities(exec)
	}

	if administrator.IsDashboard() {
		dashboard := administrator.Dashboard()
		exec, err := app.dashboardApp.Execute(dashboard, frame)
		if err != nil {
			return nil, err
		}

		builder.WithDashboard(exec)
	}

	return builder.Now()
}
