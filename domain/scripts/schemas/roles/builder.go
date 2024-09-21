package roles

import (
	"errors"
	"strings"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	read        []string
	write       []string
	review      []string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		read:        nil,
		write:       nil,
		review:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithRead adds a read to the builder
func (app *builder) WithRead(read []string) Builder {
	app.read = read
	return app
}

// WithWrite adds a write to the builder
func (app *builder) WithWrite(write []string) Builder {
	app.write = write
	return app
}

// WithReview adds a review to the builder
func (app *builder) WithReview(review []string) Builder {
	app.review = review
	return app
}

// Now builds a new Role instance
func (app *builder) Now() (Role, error) {
	if app.read != nil && len(app.read) <= 0 {
		app.read = nil
	}

	if app.write != nil && len(app.write) <= 0 {
		app.write = nil
	}

	if app.review != nil && len(app.review) <= 0 {
		app.review = nil
	}

	data := [][]byte{}
	if app.read != nil {
		data = append(data, []byte(strings.Join(app.read, ",")))
	}

	if app.write != nil {
		data = append(data, []byte(strings.Join(app.write, ",")))
	}

	if app.review != nil {
		data = append(data, []byte(strings.Join(app.review, ",")))
	}

	if len(data) <= 0 {
		return nil, errors.New("there must be at least the read, write or review in order to build a Role instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.read != nil && app.write != nil && app.review != nil {
		return createRoleWithReadAndWriteAndReview(
			*pHash,
			app.read,
			app.write,
			app.review,
		), nil
	}

	if app.read != nil && app.write != nil {
		return createRoleWithReadAndWrite(
			*pHash,
			app.read,
			app.write,
		), nil
	}

	if app.read != nil && app.review != nil {
		return createRoleWithReadAndReview(
			*pHash,
			app.read,
			app.review,
		), nil
	}

	if app.write != nil && app.review != nil {
		return createRoleWithWriteAndReview(
			*pHash,
			app.write,
			app.review,
		), nil
	}
	if app.read != nil {
		return createRoleWithRead(
			*pHash,
			app.read,
		), nil
	}

	if app.write != nil {
		return createRoleWithWrite(
			*pHash,
			app.write,
		), nil
	}

	return createRoleWithReview(*pHash, app.review), nil
}
