package administrators

import (
	"errors"

	"github.com/steve-care-software/steve/domain/accounts/administrators"
	execution_administrators "github.com/steve-care-software/steve/domain/commands/executions/visitors/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/administrators"
)

type application struct {
	adminRepository               administrators.Repository
	adminService                  administrators.Service
	adminBuilder                  administrators.Builder
	executionAdministratorBuilder execution_administrators.Builder
}

func createApplication(
	adminRepository administrators.Repository,
	adminService administrators.Service,
	adminBuilder administrators.Builder,
	executionAdministratorBuilder execution_administrators.Builder,
) Application {
	out := application{
		adminRepository:               adminRepository,
		adminService:                  adminService,
		adminBuilder:                  adminBuilder,
		executionAdministratorBuilder: executionAdministratorBuilder,
	}

	return &out
}

// Execute executes a visitor's administrator application
func (app *application) Execute(administrator inputs.Administrator) (executions.Administrator, error) {
	if administrator.IsCreate() {
		exists, err := app.adminRepository.Exists()
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, errors.New("the createAdministrator command cannot be executed when at least 1 administrator's account exists")
		}

		create := administrator.Create()
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

		return app.executionAdministratorBuilder.Create().
			WithCreate(admin).
			Now()
	}

	return nil, errors.New("the Administrator command is invalid")
}
