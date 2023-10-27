package references

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
)

type contentKeys struct {
	mp         map[uint]map[string]ContentKey
	listByKind map[uint][]ContentKey
	list       []ContentKey
}

func createContentKeys(
	mp map[uint]map[string]ContentKey,
	listByKind map[uint][]ContentKey,
	list []ContentKey,
) ContentKeys {
	out := contentKeys{
		mp:         mp,
		listByKind: listByKind,
		list:       list,
	}

	return &out
}

// List returns the contentKeys
func (obj *contentKeys) List() []ContentKey {
	return obj.list
}

// Next returns the next beginning index for a pointer
func (obj *contentKeys) Next() int64 {
	biggest := int64(0)
	for _, oneContentKey := range obj.list {
		pointer := oneContentKey.Content()
		next := int64(pointer.From() + pointer.Length())
		if biggest < next {
			biggest = next
		}
	}

	return biggest
}

// ListByKind returns the list by kind
func (obj *contentKeys) ListByKind(kind uint) ([]ContentKey, error) {
	if list, ok := obj.listByKind[kind]; ok {
		return list, nil
	}

	str := fmt.Sprintf("there is no contentKey related to the provided kind (%d)", kind)
	return nil, errors.New(str)
}

// Fetch fetches a contentKey by kind and hash
func (obj *contentKeys) Fetch(kind uint, hash hash.Hash) (ContentKey, error) {
	if mp, ok := obj.mp[kind]; ok {
		contentKeyname := hash.String()
		if ins, ok := mp[contentKeyname]; ok {
			return ins, nil
		}

		str := fmt.Sprintf("the contentKey (hash: %s) is invalid for the provided kind (%d)", contentKeyname, kind)
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("there is no contentKey related to the provided kind (%d)", kind)
	return nil, errors.New(str)
}
