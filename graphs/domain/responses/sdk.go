package responses

import "github.com/steve-care-software/steve/graphs/domain/responses/updates"

// Builder represents a response builder
type Builder interface {
	Create() Builder
	WithUpdate(update updates.Update) Builder
	Now() (Response, error)
}

// Response represents a response
type Response interface {
	HasUpdate() bool
	Update() updates.Update
}
