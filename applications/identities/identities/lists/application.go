package lists

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/lists"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/lists"
	"github.com/steve-care-software/steve/domain/stacks"
)

type application struct {
	identityRepository identities.Repository
	executionBuilder   executions.Builder
}

func createApplication(
	identityRepository identities.Repository,
	executionBuilder executions.Builder,
) Application {
	out := application{
		identityRepository: identityRepository,
		executionBuilder:   executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(list inputs.List, stack stacks.Stack) (executions.List, error) {
	usernames, err := app.identityRepository.List()
	if err != nil {
		return nil, err
	}

	variable := list.AssignToVariable()
	return app.executionBuilder.Create().
		WithVariable(variable).
		WithUsernames(usernames).
		Now()
}
