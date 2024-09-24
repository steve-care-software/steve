package rules

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Rule
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Rule) Builder {
	app.list = list
	return app
}

// Now builds a new Rules instance
func (app *builder) Now() (Rules, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Rule in order to build a Rules instance")
	}

	data := [][]byte{}
	mp := map[string]Rule{}
	for _, oneRule := range app.list {
		keyname := oneRule.Name()
		if _, ok := mp[keyname]; ok {
			str := fmt.Sprintf("the Rule (name: %s) is a duplicate", keyname)
			return nil, errors.New(str)
		}

		mp[keyname] = oneRule
		data = append(data, oneRule.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createRules(*pHash, app.list, mp), nil
}
