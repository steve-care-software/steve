package creates

import (
	"errors"

	"github.com/steve-care-software/steve/domain/accounts/visitors"
	account_visitors "github.com/steve-care-software/steve/domain/accounts/visitors"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/accounts/creates"
)

type application struct {
	visitorRepository account_visitors.Repository
	visitorService    account_visitors.Service
	visitorBuilder    account_visitors.Builder
}

func createApplication(
	visitorRepository account_visitors.Repository,
	visitorService account_visitors.Service,
	visitorBuilder account_visitors.Builder,
) Application {
	out := application{
		visitorRepository: visitorRepository,
		visitorService:    visitorService,
		visitorBuilder:    visitorBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(create inputs.Create) (visitors.Visitor, error) {
	exists, err := app.visitorRepository.Exists()
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("the createAccount command cannot be executed when at least 1 administrator's account exists")
	}

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

	return visitorAccount, nil
}
