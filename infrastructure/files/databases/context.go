package databases

import (
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/steve/domain/databases/contents"
	"github.com/steve-care-software/steve/domain/databases/references"
)

type context struct {
	identifier uint
	name       string
	pLock      *fslock.Lock
	pConn      *os.File
	reference  references.Reference
	dataOffset uint
	insertList []contents.Content
	delList    map[string]references.ContentKey
}
