package points

import "github.com/steve-care-software/steve/domain/hash"

type point struct {
	hash       hash.Hash
	name       string
	kind       uint8
	pStructure *uint8
}

func createPoint(
	hash hash.Hash,
	name string,
	kind uint8,
) Point {
	return createPointInternally(
		hash,
		name,
		kind,
		nil,
	)
}

func createPointWithStructure(
	hash hash.Hash,
	name string,
	kind uint8,
	pStructure *uint8,
) Point {
	return createPointInternally(
		hash,
		name,
		kind,
		pStructure,
	)
}

func createPointInternally(
	hash hash.Hash,
	name string,
	kind uint8,
	pStructure *uint8,
) Point {
	out := point{
		hash:       hash,
		name:       name,
		kind:       kind,
		pStructure: pStructure,
	}

	return &out
}

// Hash return the hash
func (obj *point) Hash() hash.Hash {
	return obj.hash
}

// Name return the name
func (obj *point) Name() string {
	return obj.name
}

// Kind return the kind
func (obj *point) Kind() uint8 {
	return obj.kind
}

// HasStructure returns true if there is a structure, false otherwise
func (obj *point) HasStructure() bool {
	return obj.pStructure != nil
}

// Structure returns the structure, if any
func (obj *point) Structure() *uint8 {
	return obj.pStructure
}
