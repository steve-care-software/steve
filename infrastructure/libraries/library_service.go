package libraries

import (
	"github.com/steve-care-software/steve/applications/databases"
	application_libraries "github.com/steve-care-software/steve/applications/libraries"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries"
	"github.com/steve-care-software/steve/domain/stencils/libraries/results"
	"github.com/steve-care-software/steve/domain/stencils/libraries/results/executions"
	"github.com/steve-care-software/steve/domain/stencils/libraries/results/executions/actions"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
)

type libraryService struct {
	database          databases.Application
	symbolRepository  symbols.Repository
	symbolService     symbols.Service
	resultBuilder     results.Builder
	executionsBuilder executions.Builder
	executionBuilder  executions.Builder
	actionBuilder     actions.Builder
	dbName            string
}

func createLibraryService(
	database databases.Application,
	symbolRepository symbols.Repository,
	symbolService symbols.Service,
	resultBuilder results.Builder,
	executionsBuilder executions.Builder,
	executionBuilder executions.Builder,
	actionBuilder actions.Builder,
	dbName string,
) application_libraries.Application {
	out := libraryService{
		database:          database,
		resultBuilder:     resultBuilder,
		executionsBuilder: executionsBuilder,
		executionBuilder:  executionBuilder,
		actionBuilder:     actionBuilder,
		dbName:            dbName,
	}

	return &out
}

// Save saves a library and returns the result
func (app *libraryService) Save(library libraries.Library) (results.Result, error) {
	exists, err := app.database.Exists(app.dbName)
	if err != nil {
		return nil, err
	}

	if !exists {
		err := app.database.New(app.dbName)
		if err != nil {
			return nil, err
		}
	}

	pContext, err := app.database.Open(app.dbName)
	if err != nil {
		return nil, err
	}

	defer app.database.Close(*pContext)

	path := library.Path()
	symbolsList := library.Symbols().List()
	for _, oneSymbol := range symbolsList {
		var hash hash.Hash
		symbolExists, err := app.symbolRepository.Exists(*pContext, path, hash)
		if err != nil {
			return nil, err
		}

		if symbolExists {
			continue
		}

		err = app.symbolService.Insert(*pContext, path, oneSymbol)
		if err != nil {
			return nil, err
		}

	}

	return nil, nil
}
