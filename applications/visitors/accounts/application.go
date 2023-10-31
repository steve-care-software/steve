package accounts

import (
	"errors"

	account_visitors "github.com/steve-care-software/steve/domain/accounts/visitors"
	execution_accounts "github.com/steve-care-software/steve/domain/commands/executions/visitors/accounts"
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors/accounts"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/accounts"
)

type application struct {
	visitorRepository       account_visitors.Repository
	visitorService          account_visitors.Service
	visitorBuilder          account_visitors.Builder
	executionAccountBuilder execution_accounts.Builder
}

func createApplication(
	visitorRepository account_visitors.Repository,
	visitorService account_visitors.Service,
	visitorBuilder account_visitors.Builder,
	executionAccountBuilder execution_accounts.Builder,
) Application {
	out := application{
		visitorRepository:       visitorRepository,
		visitorService:          visitorService,
		visitorBuilder:          visitorBuilder,
		executionAccountBuilder: executionAccountBuilder,
	}

	return &out
}

// Execute executes a visitor's account application
func (app *application) Execute(account inputs.Account) (executions.Account, error) {
	if account.IsCreate() {
		exists, err := app.visitorRepository.Exists()
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, errors.New("the createAccount command cannot be executed when at least 1 administrator's account exists")
		}

		create := account.Create()
		stencil := create.Stencil()
		visitorAccount, err := app.visitorBuilder.Create().
			WithStencil(stencil).
			Now()

		if err != nil {
			return nil, err
		}

		err = app.visitorService.Insert(visitorAccount)
		if err != nil {
			return nil, err
		}

		return app.executionAccountBuilder.Create().
			WithCreate(visitorAccount).
			Now()
	}

	return nil, errors.New("the Account command is invalid")
}
