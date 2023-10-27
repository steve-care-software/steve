package databases

import (
	databases "github.com/steve-care-software/steve/applications/databases"
	"github.com/steve-care-software/steve/domain/databases/contents"
	"github.com/steve-care-software/steve/domain/databases/references"
	"github.com/steve-care-software/steve/domain/trees"
)

const fileNameExtensionDelimiter = "."
const expectedReferenceBytesLength = 8
const filePermission = 0777

// NewApplication creates a new file application instance
func NewApplication(
	dirPath string,
	dstExtension string,
	bckExtension string,
	readChunkSize uint,
	onOpenFn databases.OnOpenFn,
) databases.Application {
	contentsBuilder := contents.NewBuilder()
	contentBuilder := contents.NewContentBuilder()
	referenceAdapter := references.NewAdapter()
	referenceBuilder := references.NewBuilder()
	referenceContentKeysBuilder := references.NewContentKeysBuilder()
	referenceContentKeyBuilder := references.NewContentKeyBuilder()
	referenceCommitsBuilder := references.NewCommitsBuilder()
	referenceCommitAdapter := references.NewCommitAdapter()
	referenceCommitBuilder := references.NewCommitBuilder()
	referenceActionBuilder := references.NewActionBuilder()
	referencePointerBuilder := references.NewPointerBuilder()
	hashTreeBuilder := trees.NewBuilder()
	return createApplication(
		onOpenFn,
		contentsBuilder,
		contentBuilder,
		referenceAdapter,
		referenceBuilder,
		referenceContentKeysBuilder,
		referenceContentKeyBuilder,
		referenceCommitsBuilder,
		referenceCommitAdapter,
		referenceCommitBuilder,
		referenceActionBuilder,
		referencePointerBuilder,
		hashTreeBuilder,
		dirPath,
		dstExtension,
		bckExtension,
		readChunkSize,
	)
}
