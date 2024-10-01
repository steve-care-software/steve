package parameters

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers"
)

type parameter struct {
	hash        hash.Hash
	name        string
	container   containers.Container
	isMandatory bool
}

func createParameter(
	hash hash.Hash,
	name string,
	container containers.Container,
	isMandatory bool,
) Parameter {
	out := parameter{
		hash:        hash,
		name:        name,
		container:   container,
		isMandatory: isMandatory,
	}

	return &out
}

// Hash returns the hash
func (obj *parameter) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}

// Container returns the container
func (obj *parameter) Container() containers.Container {
	return obj.container
}

// IsMandatory returns true if mandatory, false otherwise
func (obj *parameter) IsMandatory() bool {
	return obj.isMandatory
}
