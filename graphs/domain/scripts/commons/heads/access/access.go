package access

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/permissions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/writes"
)

type access struct {
	write writes.Write
	read  permissions.Permission
}

func createAccess(
	write writes.Write,
) Access {
	return createAccessInternally(write, nil)
}

func createAccessWithRead(
	write writes.Write,
	read permissions.Permission,
) Access {
	return createAccessInternally(write, read)
}

func createAccessInternally(
	write writes.Write,
	read permissions.Permission,
) Access {
	out := access{
		write: write,
		read:  read,
	}

	return &out
}

// Write returns the write
func (obj *access) Write() writes.Write {
	return obj.write
}

// HasRead returns true if there is a read, false otherwise
func (obj *access) HasRead() bool {
	return obj.read != nil
}

// Read returns the read, if any
func (obj *access) Read() permissions.Permission {
	return obj.read
}
