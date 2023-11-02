package administrators

import (
	"github.com/steve-care-software/steve/applications/administrators/administrators"
	"github.com/steve-care-software/steve/applications/administrators/identities"
	"github.com/steve-care-software/steve/applications/shares/dashboards"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators"
	"github.com/steve-care-software/steve/domain/stacks"
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
func (app *application) Execute(administrator inputs.Administrator, stack stacks.Stack) (executions.Administrator, error) {
	builder := app.executionBuilder.Create()
	if administrator.IsAdministrator() {
		admin := administrator.Administrator()
		exec, err := app.adminApp.Execute(admin, stack)
		if err != nil {
			return nil, err
		}

		builder.WithAdministrator(exec)
	}

	if administrator.IsIdentities() {
		identities := administrator.Identities()
		exec, err := app.identitiesApp.Execute(identities, stack)
		if err != nil {
			return nil, err
		}

		builder.WithIdentities(exec)
	}

	if administrator.IsDashboard() {
		dashboard := administrator.Dashboard()
		exec, err := app.dashboardApp.Execute(dashboard, stack)
		if err != nil {
			return nil, err
		}

		builder.WithDashboard(exec)
	}

	return builder.Now()
}
