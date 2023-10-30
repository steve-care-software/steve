package databases

import (
	"os"

	"github.com/juju/fslock"
)

type context struct {
	identifier uint
	name       string
	pLock      *fslock.Lock
	pConn      *os.File
	dataOffset uint
}
