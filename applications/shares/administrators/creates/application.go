package creates

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/shares/administrators/creates"
	"github.com/steve-care-software/steve/domain/commands/executions/shares/administrators/creates/failures"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/shares/administrators/creates"
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
	credentials := create.Credentials()
	username := credentials.Username()
	exists, err := app.adminRepository.Exists(username)
	if err != nil {
		return nil, err
	}

	if exists {
		username := credentials.Username()
		failure, err := app.failureBuilder.Create().
			WithUsernameAlreadyExists(username).
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()

	}

	dashboard := create.Dashboard()
	visitor := create.Visitor()
	admin, err := app.adminBuilder.Create().
		WithUsername(username).
		WithDashboard(dashboard).
		WithVisitor(visitor).
		Now()

	if err != nil {
		return nil, err
	}

	password := credentials.Password()
	err = app.adminService.Insert(admin, password)
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithSuccess(admin).
		Now()
}
