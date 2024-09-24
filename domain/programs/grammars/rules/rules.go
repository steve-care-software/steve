package rules

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
)

type rules struct {
	hash hash.Hash
	list []Rule
	mp   map[string]Rule
}

func createRules(
	hash hash.Hash,
	list []Rule,
	mp map[string]Rule,
) Rules {
	out := rules{
		hash: hash,
		list: list,
		mp:   mp,
	}

	return &out
}

// Hash returns the hash
func (obj *rules) Hash() hash.Hash {
	return obj.hash
}

// List returns the list of rule
func (obj *rules) List() []Rule {
	return obj.list
}

// Fetch fetches a rule by name
func (obj *rules) Fetch(name string) (Rule, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the rule (name: %s) does not exists", name)
	return nil, errors.New(str)
}
