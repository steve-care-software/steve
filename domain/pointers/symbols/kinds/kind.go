package kinds

import "github.com/steve-care-software/steve/domain/hash"

type kind struct {
	hash    hash.Hash
	isBytes bool
	isLayer bool
	isLink  bool
}

func createKindWithBytes(
	hash hash.Hash,
) Kind {
	return createKindInternally(hash, true, false, false)
}

func createKindWithLayer(
	hash hash.Hash,
) Kind {
	return createKindInternally(hash, false, true, false)
}

func createKindWithLink(
	hash hash.Hash,
) Kind {
	return createKindInternally(hash, false, false, true)
}

func createKindInternally(
	hash hash.Hash,
	isBytes bool,
	isLayer bool,
	isLink bool,
) Kind {
	out := kind{
		hash:    hash,
		isBytes: isBytes,
		isLayer: isLayer,
		isLink:  isLayer,
	}

	return &out
}

// Hash returns the hash
func (obj *kind) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *kind) IsBytes() bool {
	return obj.isBytes
}

// IsLayer returns true if there is layer, false otherwise
func (obj *kind) IsLayer() bool {
	return obj.isLayer
}

// IsLink returns true if there is link, false otherwise
func (obj *kind) IsLink() bool {
	return obj.isLink
}
