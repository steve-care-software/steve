package creates

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors/administrators/creates"
	"github.com/steve-care-software/steve/domain/commands/executions/visitors/administrators/creates/failures"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/administrators/creates"
)

type application struct {
	adminRepository  administrators.Repository
	adminService     administrators.Service
	adminBuilder     administrators.Builder
	failureBuilder   failures.Builder
	executionBuilder executions.Builder
}

func createApplication(
	adminRepository administrators.Repository,
	adminService administrators.Service,
	adminBuilder administrators.Builder,
	failureBuilder failures.Builder,
	executionBuilder executions.Builder,
) Application {
	out := application{
		adminRepository:  adminRepository,
		adminService:     adminService,
		adminBuilder:     adminBuilder,
		failureBuilder:   failureBuilder,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(create inputs.Create) (executions.Create, error) {
	exists, err := app.adminRepository.Exists()
	if err != nil {
		return nil, err
	}

	if exists {
		failure, err := app.failureBuilder.Create().
			AdminAlreadyExists().
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()
	}

	username := create.Username()
	dashboard := create.Dashboard()
	admin, err := app.adminBuilder.Create().
		WithUsername(username).
		WithDashboard(dashboard).
		Now()

	if err != nil {
		return nil, err
	}

	password := create.Password()
	err = app.adminService.Insert(admin, password)
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithSuccess(admin).
		Now()
}
