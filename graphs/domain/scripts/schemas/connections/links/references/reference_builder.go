package references

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
)

type referenceBuilder struct {
	internal string
	external externals.External
}

func createReferenceBuilder() ReferenceBuilder {
	out := referenceBuilder{
		internal: "",
		external: nil,
	}

	return &out
}

// Create initializes the builder
func (app *referenceBuilder) Create() ReferenceBuilder {
	return createReferenceBuilder()
}

// WithInternal adds an internal to the builder
func (app *referenceBuilder) WithInternal(internal string) ReferenceBuilder {
	app.internal = internal
	return app
}

// WithExternal adds an external to the builder
func (app *referenceBuilder) WithExternal(external externals.External) ReferenceBuilder {
	app.external = external
	return app
}

// Nwo builds a new Reference instance
func (app *referenceBuilder) Now() (Reference, error) {
	if app.internal != "" {
		return createReferenceWithInternal(app.internal), nil
	}

	if app.external != nil {
		return createReferenceWithExternal(app.external), nil
	}

	return nil, errors.New("the Reference is invalid")
}
