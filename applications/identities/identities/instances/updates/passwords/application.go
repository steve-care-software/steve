package passwords

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/credentials"
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities/identities/instances/successes/updates/passwords"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities/identities/instances/contents/updates/passwords"
)

type application struct {
	repository         identities.Repository
	service            identities.Service
	executionBuilder   executions.Builder
	credentialsBuilder credentials.Builder
}

func createApplication(
	repository identities.Repository,
	service identities.Service,
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
func (app *application) Execute(password inputs.Password, current identities.Identity) (executions.Password, error) {
	updated := password.Updated()
	username := current.Profile().Name()
	credentials, err := app.credentialsBuilder.Create().
		WithUsername(username).
		WithPassword(updated).
		Now()

	if err != nil {
		return nil, err
	}

	builder := app.executionBuilder.Create()
	currrent := password.Current()
	err = app.service.Save(current, currrent, updated)
	if err != nil {
		return builder.WithFailure(credentials).
			Now()
	}

	return builder.WithSuccess(credentials).
		Now()
}
