package applications

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
)

// Application represents the application
type Application interface {
	Begin() (*uint, error)
	Init(context uint, root roots.Root, path string) error
	Source(context uint, path string) error
	Execute(context uint, input []byte, frame frames.Frame) ([]byte, error)
	Queue(context uint) (commands.Commands, error)
	Commit(context uint, message string) error
	Back(context uint) error
	Clear(context uint) error
	Rollback(context uint) error
	Reset(context uint) error
}
