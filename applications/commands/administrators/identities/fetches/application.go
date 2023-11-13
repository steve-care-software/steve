package fetches

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/identities/successes/fetches"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/identities/successes/fetches/failures"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/identities/successes/fetches/successes"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/identities/contents/fetches"
)

type application struct {
	executionBuilder executions.Builder
	successBuilder   successes.Builder
	failureBuilder   failures.Builder
}

// Execute executes the application
func (app *application) Execute(instance inputs.Fetch, current identities.Identities) (executions.Fetch, error) {
	property := instance.Property()
	successBuilder := app.successBuilder.Create()
	if property.IsAmount() {
		list := current.List()
		amount := uint(len(list))
		successBuilder.WithAmount(amount)
	}

	if property.IsAtIndex() {
		pAtIndex := property.AtIndex()
		if current.Exceeds(*pAtIndex) {
			failure, err := app.failureBuilder.Create().
				IsIndexExceedAmount().
				WithIndex(*pAtIndex).
				Now()

			if err != nil {
				return nil, err
			}

			return app.executionBuilder.Create().
				WithFailure(failure).
				Now()
		}

		identity, err := current.Fetch(*pAtIndex)
		if err != nil {
			return nil, err
		}

		successBuilder.WithAtIndex(identity)

	}

	success, err := successBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().
		WithSuccess(success).
		Now()
}
