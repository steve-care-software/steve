package transactions

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Transaction
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
func (app *builder) WithList(list []Transaction) Builder {
	app.list = list
	return app
}

// Now builds a new Transactions instance
func (app *builder) Now() (Transactions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Transaction in order to build a Transactions instance")
	}

	scriptMap := map[string]Transaction{}
	for _, oneTrx := range app.list {
		keyname := string(oneTrx.Entry().Script())
		if _, ok := scriptMap[keyname]; ok {
			str := fmt.Sprintf("the script (%s) has already been added in a previous transaction", keyname)
			return nil, errors.New(str)
		}

		scriptMap[keyname] = oneTrx
	}

	data := [][]byte{}
	for _, oneTrx := range app.list {
		data = append(data, oneTrx.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createTransactions(*pHash, app.list), nil
}
