package modifications

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources"
)

type modificationBuilder struct {
	hashAdapter hash.Adapter
	insert      resources.Resource
	save        resources.Resource
	delete      string
}

func createModificationBuilder(
	hashAdapter hash.Adapter,
) ModificationBuilder {
	out := modificationBuilder{
		hashAdapter: hashAdapter,
		insert:      nil,
		save:        nil,
		delete:      "",
	}

	return &out
}

// Create initializes the builder
func (app *modificationBuilder) Create() ModificationBuilder {
	return createModificationBuilder(
		app.hashAdapter,
	)
}

// WithInsert adds an insert to the builder
func (app *modificationBuilder) WithInsert(insert resources.Resource) ModificationBuilder {
	app.insert = insert
	return app
}

// WithSave adds a save to the builder
func (app *modificationBuilder) WithSave(save resources.Resource) ModificationBuilder {
	app.save = save
	return app
}

// WithDelete adds a delete to the builder
func (app *modificationBuilder) WithDelete(delete string) ModificationBuilder {
	app.delete = delete
	return app
}

// Now builds a new Modification instance
func (app *modificationBuilder) Now() (Modification, error) {
	data := [][]byte{}
	if app.insert != nil {
		data = append(data, app.insert.Hash().Bytes())
	}

	if app.save != nil {
		data = append(data, app.save.Hash().Bytes())
	}

	if app.delete != "" {
		data = append(data, []byte(app.delete))
	}

	if len(data) != 1 {
		return nil, errors.New("the Modification is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil {
		return createModificationWithInsert(*pHash, app.insert), nil
	}

	if app.save != nil {
		return createModificationWithSave(*pHash, app.save), nil
	}

	return createModificationWithDelete(*pHash, app.delete), nil
}
