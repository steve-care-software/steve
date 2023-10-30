package blocks

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/votes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/pointers"
	"github.com/steve-care-software/steve/domain/hash"
)

// Builder represents a block builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithVote(vote votes.Vote) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Content() Content
	Vote() votes.Vote
	Pointer() pointers.Pointer
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithCommits(commits commits.Commits) ContentBuilder
	WithParent(parent Parent) ContentBuilder
	Now() (Content, error)
}

// Content represents a block content
type Content interface {
	Hash() hash.Hash
	Commits() commits.Commits
	Parent() Parent
	Pointer() pointers.Pointer
}

// ParentBuilder represents a parent builder
type ParentBuilder interface {
	Create() ParentBuilder
	WithBlock(block hash.Hash) ParentBuilder
	WithRoot(root resources.Resources) ParentBuilder
	Now() (Parent, error)
}

// Parent represents a parent block
type Parent interface {
	Hash() hash.Hash
	IsBlock() bool
	Block() Block
	IsRoot() bool
	Root() resources.Resources
}
