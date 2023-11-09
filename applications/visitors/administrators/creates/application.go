package creates

import (
	"errors"

	"github.com/steve-care-software/steve/domain/accounts/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/administrators/creates"
)

type application struct {
	adminRepository administrators.Repository
	adminService    administrators.Service
	adminBuilder    administrators.Builder
}

func createApplication(
	adminRepository administrators.Repository,
	adminService administrators.Service,
	adminBuilder administrators.Builder,
) Application {
	out := application{
		adminRepository: adminRepository,
		adminService:    adminService,
		adminBuilder:    adminBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(create inputs.Create) (administrators.Administrator, error) {
	exists, err := app.adminRepository.Exists()
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("the createAdministrator command cannot be executed when at least 1 administrator's account exists")
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

	return admin, nil
}
