package contexts

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/connections/links/contexts"
)

// Application represents the context application
type Application interface {
	Children(identifier uuid.UUID) (contexts.Contexts, error)
	ChildrenByPath(path []string) (contexts.Contexts, error)
	Retrieve(identifier uuid.UUID) (contexts.Context, error)
	RetrieveByPath(path []string) (contexts.Context, error)
	Insert(context contexts.Context) error
	Update(identifier uuid.UUID, newName string) error
	Move(identifier uuid.UUID, newParent uuid.UUID) error
	MoveToPath(identifier uuid.UUID, toPath []string) error
	Delete(identifier uuid.UUID) error
}
