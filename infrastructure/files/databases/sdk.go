package databases

import (
	databases "github.com/steve-care-software/steve/applications/blockchains/databases"
)

const fileNameExtensionDelimiter = "."
const filePermission = 0777

// NewApplication creates a new file application instance
func NewApplication(
	dirPath string,
	dstExtension string,
	bckExtension string,
	onOpenFn databases.OnOpenFn,
) databases.Application {
	return createApplication(
		onOpenFn,
		dirPath,
		dstExtension,
		bckExtension,
	)
}
