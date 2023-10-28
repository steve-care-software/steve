package applications

import (
	"errors"
	"fmt"

	application_layers "github.com/steve-care-software/steve/applications/layers"
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/signatures"
	"github.com/steve-care-software/steve/domain/stencils"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
	"github.com/steve-care-software/steve/domain/stencils/messages"
	"github.com/steve-care-software/steve/domain/stencils/queries"
	"github.com/steve-care-software/steve/domain/stencils/results/executions"
)

type application struct {
	layerApp         application_layers.Application
	adminRepository  administrators.Repository
	adminService     administrators.Service
	symbolRepository symbols.Repository
	queryBuilder     queries.Builder
}

func createApplication() Application {
	out := application{}
	return &out
}

// Authorize executes an authorized query
func (app *application) Authorize(message messages.Message, username string, password []byte) (executions.Execution, error) {
	adminIns, err := app.adminRepository.Retrieve(username, password)
	if err != nil {
		return nil, err
	}

	container := []string{}
	symbolHash := adminIns.Dashboard().Root().Root()
	symbol, err := app.symbolRepository.Retrieve(container, symbolHash)
	if err != nil {
		return nil, err
	}

	if !symbol.IsLayer() {
		str := fmt.Sprintf("the Symbol (hash: %s) was expected to contain a Layer", symbolHash.String())
		return nil, errors.New(str)
	}

	layer := symbol.Layer()
	query, err := app.queryBuilder.Create().
		WithMessage(message).
		WithLayer(layer).
		Now()

	if err != nil {
		return nil, err
	}

	result, err := app.layerApp.Execute(query)
	if err != nil {
		return nil, err
	}

	// find the link to execute, then execute it.

	fmt.Printf("\n%v\n", result)

	return nil, nil
}

// Authenticate executes an authenticated query
func (app *application) Authenticate(message messages.Message, signature signatures.Signature) (executions.Execution, error) {
	return nil, nil
}

// Visit executes a visitor query
func (app *application) Visit(message messages.Message, stencil stencils.Stencil) (executions.Execution, error) {
	return nil, nil
}
