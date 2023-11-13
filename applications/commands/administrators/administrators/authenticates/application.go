package authenticates

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/authenticates"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/authenticates/successes"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/administrators/authenticates"
)

type application struct {
	adminRepository  administrators.Repository
	executionBuilder executions.Builder
	successBuilder   successes.Builder
}

func createApplication(
	adminRepository administrators.Repository,
	executionBuilder executions.Builder,
	successBuilder successes.Builder,
) Application {
	out := application{
		adminRepository:  adminRepository,
		executionBuilder: executionBuilder,
		successBuilder:   successBuilder,
	}

	return &out
}

// Execute executes the authentication
func (app *application) Execute(administrator inputs.Authenticate) (executions.Authenticate, error) {
	credentials := administrator.Credentials()
	admin, err := app.adminRepository.Retrieve(credentials)
	builder := app.executionBuilder.Create()
	if err != nil {
		return builder.WithFailure(credentials).
			Now()
	}

	variable := administrator.AssignToVariable()
	success, err := app.successBuilder.Create().
		WithVariable(variable).
		WithInstance(admin).
		Now()

	if err != nil {
		return nil, err
	}

	return builder.WithSuccess(success).
		Now()

}
