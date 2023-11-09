package encryptors

import (
	"github.com/steve-care-software/steve/domain/commands/inputs/identities/encryptors/contents"
)

// Builder represents the encryptor builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithAssignToVariable(assignToVariable string) Builder
	WithContent(content contents.Content) Builder
	Now() (Encryptor, error)
}

// Encryptor represents an encryptor
type Encryptor interface {
	Name() string
	AssignToVariable() string
	Content() contents.Content
}
