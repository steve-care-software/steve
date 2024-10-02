package constants

import (
	"errors"
	"fmt"
)

type constants struct {
	list []Constant
	mp   map[string]Constant
}

func createConstants(
	list []Constant,
	mp map[string]Constant,
) Constants {
	out := constants{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of constant
func (obj *constants) List() []Constant {
	return obj.list
}

// Fetch fetches a constant by name
func (obj *constants) Fetch(name string) (Constant, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the constant (name: %s) does not exists", name)
	return nil, errors.New(str)
}
