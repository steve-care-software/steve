package authenticates

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators/authenticates"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/authenticates"
)

type application struct {
	adminRepository  administrators.Repository
	executionBuilder executions.Builder
}

func createApplication(
	adminRepository administrators.Repository,
	executionBuilder executions.Builder,
) Application {
	out := application{
		adminRepository:  adminRepository,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the authentication
func (app *application) Execute(administrator inputs.Authenticate) (executions.Authenticate, error) {
	username := administrator.Username()
	password := administrator.Password()
	admin, err := app.adminRepository.Retrieve(username, password)
	if err != nil {
		return nil, err
	}

	variable := administrator.AssignToVariable()
	return app.executionBuilder.Create().
		WithVariable(variable).
		WithInstance(admin).
		Now()
}
