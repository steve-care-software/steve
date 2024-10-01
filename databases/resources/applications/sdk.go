package resources

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/databases/resources/domain/contents"
	"github.com/steve-care-software/steve/databases/resources/domain/headers"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

const appNotInitErr = "the application has not been initialized yet"
const alreadyModifiedErrPattern = "the resource (identifier: %s) has already been modified in the current session"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	contentBuilder := contents.NewBuilder()
	headerBuilder := headers.NewBuilder()
	activityBuilder := activities.NewBuilder()
	commitsBuilder := commits.NewBuilder()
	commitBuilder := commits.NewCommitBuilder()
	modificationsBuilder := modifications.NewBuilder()
	modificationBuilder := modifications.NewModificationBuilder()
	resourcesBuilder := resources.NewBuilder()
	resourceBuilder := resources.NewResourceBuilder()
	pointerBuilder := pointers.NewBuilder()
	headerAdapter := headers.NewAdapter()
	return createBuilder(
		contentBuilder,
		headerBuilder,
		activityBuilder,
		commitsBuilder,
		commitBuilder,
		modificationsBuilder,
		modificationBuilder,
		resourcesBuilder,
		resourceBuilder,
		pointerBuilder,
		headerAdapter,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithReadChunkSize(readChkSize uint64) Builder
	WithBasePath(basePath string) Builder
	WithTargetIdentifier(targetIdentifier string) Builder
	Now() (Application, error)
}

// Application represents the resources application
type Application interface {
	Init(dbIdentifier string) error
	Retrieve(identifier string) ([]byte, error)
	Insert(identifier string, data []byte) error
	Save(identifier string, data []byte) error
	Delete(identifier string) error
	Head() (commits.Commit, error)
	Commit() error
	Cancel() error
	Rollback(amount uint) error
	RollbackTo(commit hash.Hash) error
}
