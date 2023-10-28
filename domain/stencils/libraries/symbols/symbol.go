package symbols

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links"
)

type symbol struct {
	hash  hash.Hash
	bytes []byte
	layer layers.Layer
	link  links.Link
}

func createSymbolWithBytes(
	hash hash.Hash,
	bytes []byte,
) Symbol {
	return createSymbolInternally(
		hash,
		bytes,
		nil,
		nil,
	)
}

func createSymbolWithLayer(
	hash hash.Hash,
	layer layers.Layer,
) Symbol {
	return createSymbolInternally(
		hash,
		nil,
		layer,
		nil,
	)
}

func createSymbolWithLink(
	hash hash.Hash,
	link links.Link,
) Symbol {
	return createSymbolInternally(
		hash,
		nil,
		nil,
		link,
	)
}

func createSymbolInternally(
	hash hash.Hash,
	bytes []byte,
	layer layers.Layer,
	link links.Link,
) Symbol {
	out := symbol{
		hash:  hash,
		bytes: bytes,
		layer: layer,
		link:  link,
	}

	return &out
}

// Hash returns the hash
func (obj *symbol) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *symbol) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *symbol) Bytes() []byte {
	return obj.bytes
}

// IsLayer returns true if there is a layer, false otherwise
func (obj *symbol) IsLayer() bool {
	return obj.layer != nil
}

// Layer returns a layer, if any
func (obj *symbol) Layer() layers.Layer {
	return obj.layer
}

// IsLink returns true if there is a link, false otherwise
func (obj *symbol) IsLink() bool {
	return obj.link != nil
}

// Link returns a link, if any
func (obj *symbol) Link() links.Link {
	return obj.link
}
