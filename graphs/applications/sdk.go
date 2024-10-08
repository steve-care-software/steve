package applications

import (
	"github.com/steve-care-software/steve/graphs/domain/responses"
	"github.com/steve-care-software/steve/graphs/domain/scripts"
)

// Application represents the graphdb application
type Application interface {
	Execute(script scripts.Script) (responses.Response, error)
}
