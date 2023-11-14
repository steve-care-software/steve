package frames

import (
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/scopes/assignments/frames"
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
	List() []Entry
	Fetch(index uint) (Entry, error)
}

// EntryBuilder represents an entry builder
type EntryBuilder interface {
	Create() EntryBuilder
	IsStop() EntryBuilder
	WithIndex(index uint) EntryBuilder
	WithFrame(frame frames.Frame) EntryBuilder
	WithBlock(block Block) EntryBuilder
	Now() (Entry, error)
}

// Entry represents a frame entry
type Entry interface {
	Hash() hash.Hash
	Index() uint
	Content() Content
}

// Content represents an entry content
type Content interface {
	Hash() hash.Hash
	IsStop() bool
	IsFrame() bool
	Frame() frames.Frame
	IsBlock() bool
	Block() Block
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithCondition(condition []byte) BlockBuilder
	WithFrames(frames frames.Frames) BlockBuilder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Condition() []byte
	Frames() frames.Frames
}
