package references

import (
	"errors"
	"fmt"
)

type contentContentKeysBuilder struct {
	list []ContentKey
}

func createContentKeysBuilder() ContentKeysBuilder {
	out := contentContentKeysBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentContentKeysBuilder) Create() ContentKeysBuilder {
	return createContentKeysBuilder()
}

// WithList add contentContentKeys to the builder
func (app *contentContentKeysBuilder) WithList(list []ContentKey) ContentKeysBuilder {
	app.list = list
	return app
}

// Now builds ContentKeys instance
func (app *contentContentKeysBuilder) Now() (ContentKeys, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ContentKey in order to build a ContentKeys instance")
	}

	mp := map[uint]map[string]ContentKey{}
	listByKind := map[uint][]ContentKey{}
	for _, oneContentKey := range app.list {
		kind := oneContentKey.Kind()
		if _, ok := mp[kind]; !ok {
			mp[kind] = map[string]ContentKey{}
		}

		if _, ok := listByKind[kind]; !ok {
			listByKind[kind] = []ContentKey{}
		}

		keyname := oneContentKey.Hash().String()
		if _, ok := mp[kind][keyname]; ok {
			str := fmt.Sprintf("this resource (kind: %d, hash: %s) already exists", kind, keyname)
			return nil, errors.New(str)
		}

		mp[kind][keyname] = oneContentKey
		listByKind[kind] = append(listByKind[kind], oneContentKey)
	}

	return createContentKeys(mp, listByKind, app.list), nil
}
