package administrators

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/administrators"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes an administrator
func (app *application) Execute(administrator inputs.Administrator) (executions.Administrator, error) {
	if administrator.IsAdministrator() {

	}

	if administrator.IsIdentities() {

	}

	if administrator.IsDashboard() {

	}

	return nil, nil
}
