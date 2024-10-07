package walkers

import (
	"errors"
	"fmt"
)

type tokenlist struct {
	fn   MapFn
	list []SelectedTokenList
	mp   map[string]SelectedTokenList
}

func createTokenList(
	fn MapFn,
	list []SelectedTokenList,
	mp map[string]SelectedTokenList,
) TokenList {
	out := tokenlist{
		fn:   fn,
		list: list,
		mp:   mp,
	}

	return &out
}

// Fn returns the MapFn
func (obj *tokenlist) Fn() MapFn {
	return obj.fn
}

// List returns the list
func (obj *tokenlist) List() []SelectedTokenList {
	return obj.list
}

// Fetch fetches a SelectedTokenList by name
func (obj *tokenlist) Fetch(name string) (SelectedTokenList, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the SelectedTokenList (name: %s) does not exists", name)
	return nil, errors.New(str)
}
