package frames

import (
	bytes_frames "github.com/steve-care-software/steve/vms/children/bytes/frames"
	assignable_results "github.com/steve-care-software/steve/vms/children/layers/scopes/instructions/scopes/assignments/scopes/assignables/results"
	"github.com/steve-care-software/steve/vms/libraries/hash"
)

// Builder represents the frames builder
type Builder interface {
	Create() Builder
	WithList(list []Frame) Builder
	Now() (Frames, error)
}

// Frames represents frames
type Frames interface {
	Hash() hash.Hash
	List() []Frame
}

// FrameBuilder represents a frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WithList(list []Assignment) FrameBuilder
	Now() (Frame, error)
}

// Frame represents a frame
type Frame interface {
	Hash() hash.Hash
	Bytes() bytes_frames.Frame
	Fetch(name string) (assignable_results.Result, error)
	List() []Assignment
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithValue(value assignable_results.Result) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Hash() hash.Hash
	Name() string
	Value() assignable_results.Result
}
