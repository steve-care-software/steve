package points

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/points"
)

// Application represents the point application
type Application interface {
	RetrieveByConnection(connIdentifier uuid.UUID) (points.Point, error)
	Insert(point points.Point) error
	Delete(identifier uuid.UUID) error
}
