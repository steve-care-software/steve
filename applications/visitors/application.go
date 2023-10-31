package visitors

import (
	"errors"

	"github.com/steve-care-software/steve/domain/accounts/administrators"
	account_visitors "github.com/steve-care-software/steve/domain/accounts/visitors"
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors"
	execution_accounts "github.com/steve-care-software/steve/domain/commands/executions/visitors/accounts"
	execution_administrators "github.com/steve-care-software/steve/domain/commands/executions/visitors/administrators"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors"
)

type application struct {
	adminRepository               administrators.Repository
	adminService                  administrators.Service
	adminBuilder                  administrators.Builder
	visitorRepository             account_visitors.Repository
	visitorService                account_visitors.Service
	visitorBuilder                account_visitors.Builder
	executionBuilder              executions.Builder
	executionAccountBuilder       execution_accounts.Builder
	executionAdministratorBuilder execution_administrators.Builder
}

func createApplication(
	adminRepository administrators.Repository,
	adminService administrators.Service,
	adminBuilder administrators.Builder,
	visitorRepository account_visitors.Repository,
	visitorService account_visitors.Service,
	visitorBuilder account_visitors.Builder,
	executionAccountBuilder execution_accounts.Builder,
	executionAdministratorBuilder execution_administrators.Builder,
) Application {
	out := application{
		adminRepository:               adminRepository,
		adminService:                  adminService,
		adminBuilder:                  adminBuilder,
		visitorRepository:             visitorRepository,
		visitorService:                visitorService,
		visitorBuilder:                visitorBuilder,
		executionAccountBuilder:       executionAccountBuilder,
		executionAdministratorBuilder: executionAdministratorBuilder,
	}

	return &out
}

// Execute executes a visitor's application
func (app *application) Execute(visitor inputs.Visitor) (executions.Visitor, error) {
	if visitor.IsAdministrator() {
		cmdAdmin := visitor.Administrator()
		if cmdAdmin.IsCreate() {
			exists, err := app.adminRepository.Exists()
			if err != nil {
				return nil, err
			}

			if exists {
				return nil, errors.New("the createAdministrator command cannot be executed when at least 1 administrator's account exists")
			}

			create := cmdAdmin.Create()
			username := create.Username()
			dashboard := create.Dashboard()
			admin, err := app.adminBuilder.Create().
				WithUsername(username).
				WithDashboard(dashboard).
				Now()

			if err != nil {
				return nil, err
			}

			password := create.Password()
			err = app.adminService.Insert(admin, password)
			if err != nil {
				return nil, err
			}

			execCreate, err := app.executionAdministratorBuilder.Create().
				WithCreate(admin).
				Now()

			if err != nil {
				return nil, err
			}

			return app.executionBuilder.Create().
				WithAdministrator(execCreate).
				Now()
		}

		return nil, errors.New("the Administrator command is invalid")
	}

	if visitor.IsAccount() {
		account := visitor.Account()
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

			execAccount, err := app.executionAccountBuilder.Create().
				WithCreate(visitorAccount).
				Now()

			if err != nil {
				return nil, err
			}

			return app.executionBuilder.Create().
				WithAccount(execAccount).
				Now()
		}

		return nil, errors.New("the Account command is invalid")

	}
	return nil, nil
}
