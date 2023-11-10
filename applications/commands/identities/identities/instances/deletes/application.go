package deletes

import (
	"github.com/steve-care-software/steve/domain/accounts/credentials"
	"github.com/steve-care-software/steve/domain/accounts/identities"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities/instances/successes/deletes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/identities/identities/instances/successes/deletes/failures"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/identities/instances/contents/deletes"
)

type application struct {
	identityRepository identities.Repository
	identityService    identities.Service
	executionBuilder   executions.Builder
	failureBuilder     failures.Builder
	credentialsBuilder credentials.Builder
}

func createApplication(
	identityRepository identities.Repository,
	identityService identities.Service,
	executionBuilder executions.Builder,
	failureBuilder failures.Builder,
	credentialsBuilder credentials.Builder,
) Application {
	out := application{
		identityRepository: identityRepository,
		executionBuilder:   executionBuilder,
		failureBuilder:     failureBuilder,
		credentialsBuilder: credentialsBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(delete inputs.Delete, current identities.Identity) (executions.Delete, error) {
	password := delete.Password()
	username := current.Profile().Name()
	credentials, err := app.credentialsBuilder.Create().
		WithUsername(username).
		WithPassword(password).
		Now()

	if err != nil {
		return nil, err
	}

	retIdentity, err := app.identityRepository.Retrieve(credentials)
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

	err = app.identityService.Delete(credentials)
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
		WithSuccess(retIdentity).
		Now()
}
