package authenticates

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/authenticates"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/identities/authenticates/successes"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/authenticates"
)

type application struct {
	identityRepository identities.Repository
	executionBuilder   executions.Builder
	successBuilder     successes.Builder
}

func createApplication(
	identityRepository identities.Repository,
	executionBuilder executions.Builder,
	successBuilder successes.Builder,
) Application {
	out := application{
		identityRepository: identityRepository,
		executionBuilder:   executionBuilder,
		successBuilder:     successBuilder,
	}

	return &out
}

// Execute executes the authentication
func (app *application) Execute(authenticate inputs.Authenticate) (executions.Authenticate, error) {
	credentials := authenticate.Credentials()
	identity, err := app.identityRepository.Retrieve(credentials)
	builder := app.executionBuilder.Create()
	if err != nil {
		return builder.WithFailure(credentials).
			Now()
	}

	variable := authenticate.AssignToVariable()
	success, err := app.successBuilder.Create().
		WithVariable(variable).
		WithInstance(identity).
		Now()

	if err != nil {
		return nil, err
	}

	return builder.WithSuccess(success).
		Now()

}
