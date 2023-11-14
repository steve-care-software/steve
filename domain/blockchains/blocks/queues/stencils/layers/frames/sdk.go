package frames

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/stencils/layers/frames/assignables"

// Frames represents frames
type Frames interface {
	List() []Frame
}

// Frame represents a frame
type Frame interface {
	Fetch(name string) (assignables.Assignable, error)
}
