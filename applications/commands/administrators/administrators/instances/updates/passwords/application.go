package passwords

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/accounts/credentials"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/updates/passwords"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/administrators/instances/contents/updates/passwords"
)

type application struct {
	repository         administrators.Repository
	service            administrators.Service
	executionBuilder   executions.Builder
	credentialsBuilder credentials.Builder
}

func createApplication(
	repository administrators.Repository,
	service administrators.Service,
	credentialsBuilder credentials.Builder,
) Application {
	out := application{
		repository:         repository,
		service:            service,
		credentialsBuilder: credentialsBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(password inputs.Password, instance administrators.Administrator) (executions.Password, error) {
	updated := password.Updated()
	username := instance.Username()
	credentials, err := app.credentialsBuilder.Create().
		WithUsername(username).
		WithPassword(updated).
		Now()

	if err != nil {
		return nil, err
	}

	builder := app.executionBuilder.Create()
	currrent := password.Current()
	err = app.service.Save(instance, currrent, updated)
	if err != nil {
		return builder.WithFailure(credentials).
			Now()
	}

	return builder.WithSuccess(credentials).
		Now()
}
