package applications

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
)

// Application represents the application
type Application interface {
	Init(root roots.Root, path string) error
	Begin(path string) (*uint, error)
	Exists(context uint) bool
	Execute(context uint, input []byte, frame frames.Frame) ([]byte, error)
	Queue(context uint) (queues.Queue, error)
	Commit(context uint, message string) error
	Back(context uint) error
	Clear(context uint) error
	Rollback(context uint) error
}
