package chains

import "github.com/steve-care-software/steve/domain/chains"

// Application represents the chain application
type Application interface {
	Execute(chain chains.Chain) ([]byte, error)
}
