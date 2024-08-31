package contexts

import "github.com/google/uuid"

type context struct {
	identifier uuid.UUID
	name       string
	pParent    *uuid.UUID
}

func createContext(
	identifier uuid.UUID,
	name string,
) Context {
	return createContextInternally(identifier, name, nil)
}

func createContextWithParent(
	identifier uuid.UUID,
	name string,
	pParent *uuid.UUID,
) Context {
	return createContextInternally(identifier, name, pParent)
}

func createContextInternally(
	identifier uuid.UUID,
	name string,
	pParent *uuid.UUID,
) Context {
	out := context{
		identifier: identifier,
		name:       name,
		pParent:    pParent,
	}

	return &out
}

// Identifier returns the identifier
func (obj *context) Identifier() uuid.UUID {
	return obj.identifier
}

// Name returns the name
func (obj *context) Name() string {
	return obj.name
}

// HasParent returns true if there is parent, false otherwise
func (obj *context) HasParent() bool {
	return obj.pParent != nil
}

// Parent returns the parent, if any
func (obj *context) Parent() *uuid.UUID {
	return obj.pParent
}
