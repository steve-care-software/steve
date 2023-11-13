package signers

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/signers/contents"

// Builder represents the signer builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithAssignToVariable(assignToVariable string) Builder
	WithContent(content contents.Content) Builder
	Now() (Signer, error)
}

// Signer represents a signer
type Signer interface {
	Name() string
	AssignToVariable() string
	Content() contents.Content
}
