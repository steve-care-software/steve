package resources

import (
	"github.com/steve-care-software/steve/domain/stores/commits/modifications/resources"
	"github.com/steve-care-software/steve/domain/stores/commits/modifications/resources/pointers"
)

type application struct {
	resourcesAdapter resources.Adapter
	resourcesBuilder resources.Builder
	resourceBuilder  resources.ResourceBuilder
	pointerAdapter   pointers.Adapter
	pointerBuilder   pointers.Builder
}

func createApplication(
	resourcesAdapter resources.Adapter,
	resourcesBuilder resources.Builder,
	resourceBuilder resources.ResourceBuilder,
	pointerAdapter pointers.Adapter,
	pointerBuilder pointers.Builder,
) Application {
	out := application{
		resourcesAdapter: resourcesAdapter,
		resourcesBuilder: resourcesBuilder,
		resourceBuilder:  resourceBuilder,
		pointerAdapter:   pointerAdapter,
		pointerBuilder:   pointerBuilder,
	}

	return &out
}

// Begin begins the modifications
func (obj *application) Begin(dbIdentifier string) error {
	return nil
}

// Retrieve retrieves bytes from an identifier
func (app *application) Retrieve(identifier string) ([]byte, error) {
	return nil, nil
}

// Insert inserts data into an identifier
func (app *application) Insert(identifier string, data []byte) error {
	return nil
}

// Save saves data into an identifier
func (app *application) Save(identifier string, data []byte) error {
	return nil
}

// Delete deletes an identifier
func (app *application) Delete(identifier string) error {
	return nil
}

// Commit commits the modifications
func (obj *application) Commit() error {
	return nil
}

// Cancel cancels the modifications
func (obj *application) Cancel() error {
	return nil
}

// Rollback remove the last commits
func (obj *application) Rollback(amount uint) error {
	return nil
}
