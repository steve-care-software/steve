package frames

import (
	bytes_frames "github.com/steve-care-software/steve/vms/bytes/frames"
	bytes_results "github.com/steve-care-software/steve/vms/bytes/results"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	assignment_frames "github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/frames"
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
	WithList(list []Entry) FrameBuilder
	Now() (Frame, error)
}

// Frame represents a frame
type Frame interface {
	Hash() hash.Hash
	Bytes() bytes_frames.Frame
	Assignment() assignment_frames.Frame
	List() []Entry
	Fetch(index uint) (Entry, error)
}

// EntryBuilder represents an entry builder
type EntryBuilder interface {
	Create() EntryBuilder
	WithFrame(frame assignment_frames.Frame) EntryBuilder
	WithBlock(block Block) EntryBuilder
	Now() (Entry, error)
}

// Entry represents a frame entry
type Entry interface {
	Hash() hash.Hash
	IsFrame() bool
	Frame() assignment_frames.Frame
	IsBlock() bool
	Block() Block
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithSuccess(success Frames) BlockBuilder
	WithFailure(failure BlockFailure) BlockBuilder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	IsSuccess() bool
	Success() Frames
	IsFailure() bool
	Failure() BlockFailure
}

// BlockFailureBuilder represents a block failure builder
type BlockFailureBuilder interface {
	Create() BlockFailureBuilder
	WithConditionFailed(conditionFailed bytes_results.Result) BlockFailureBuilder
	Now() (BlockFailure, error)
}

// BlockFailure represents a block failure
type BlockFailure interface {
	IsConditionFailed() bool
	ConditionFailed() bytes_results.Result
}
