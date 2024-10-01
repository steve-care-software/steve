package chains

import "github.com/steve-care-software/steve/engine/domain/chains"

// Application represents the chain application
type Application interface {
	Execute(chain chains.Chain, input []byte) ([]byte, error)
}
