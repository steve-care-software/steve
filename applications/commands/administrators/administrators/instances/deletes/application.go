package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/accounts/credentials"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/deletes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/deletes/failures"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/administrators/instances/contents/deletes"
)

type application struct {
	adminRepository    administrators.Repository
	adminService       administrators.Service
	executionBuilder   executions.Builder
	failureBuilder     failures.Builder
	credentialsBuilder credentials.Builder
}

func createApplication(
	adminRepository administrators.Repository,
	adminService administrators.Service,
	executionBuilder executions.Builder,
	failureBuilder failures.Builder,
	credentialsBuilder credentials.Builder,
) Application {
	out := application{
		adminRepository:    adminRepository,
		executionBuilder:   executionBuilder,
		failureBuilder:     failureBuilder,
		credentialsBuilder: credentialsBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(delete inputs.Delete, current administrators.Administrator) (executions.Delete, error) {
	password := delete.Password()
	username := current.Username()
	credentials, err := app.credentialsBuilder.Create().
		WithUsername(username).
		WithPassword(password).
		Now()

	if err != nil {
		return nil, err
	}

	retAdmin, err := app.adminRepository.Retrieve(credentials)
	if err != nil {
		failure, err := app.failureBuilder.Create().
			CouldNotRetrieve().
			WithCredentials(credentials).
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()
	}

	err = app.adminService.Delete(credentials)
	if err != nil {
		failure, err := app.failureBuilder.Create().
			CouldNotDelete().
			WithCredentials(credentials).
			Now()

		if err != nil {
			return nil, err
		}

		return app.executionBuilder.Create().
			WithFailure(failure).
			Now()
	}

	return app.executionBuilder.Create().
		WithSuccess(retAdmin).
		Now()
}
