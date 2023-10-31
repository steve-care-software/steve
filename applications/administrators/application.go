package administrators

import (
	"github.com/steve-care-software/steve/domain/commands/administrators/executions"
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs"
)

type application struct {
	inputAdapter inputs.Adapter
}

func createApplication(
	inputAdapter inputs.Adapter,
) Application {
	out := application{
		inputAdapter: inputAdapter,
	}

	return &out
}

// Execute executes a command
func (app *application) Execute(message []byte, username string, password []byte) (executions.Execution, error) {
	input, err := app.inputAdapter.ToInput(message)
	if err != nil {
		return nil, err
	}

	if input.IsAdministrator() {

	}

	if input.IsIdentities() {

	}

	if input.IsDashboard() {

	}

	return nil, nil
}
