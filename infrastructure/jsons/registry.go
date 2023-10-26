package jsons

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
)

type registry struct {
	symbols map[string]symbols.Symbol
}

func createRegistry() *registry {
	out := registry{
		symbols: nil,
	}

	ptr := &out
	return ptr.init()
}

func (app *registry) init() *registry {
	app.symbols = map[string]symbols.Symbol{}
	return app
}

func (app *registry) register(name string, symbol symbols.Symbol) error {
	if _, ok := app.symbols[name]; ok {
		str := fmt.Sprintf("the symbol (name: %s) already exists", name)
		return errors.New(str)
	}

	app.symbols[name] = symbol
	return nil
}

func (app *registry) retrieve(name string) (symbols.Symbol, error) {
	if _, ok := app.symbols[name]; !ok {
		str := fmt.Sprintf("the symbol (name: %s) does NOT exists", name)
		return nil, errors.New(str)
	}

	return app.symbols[name], nil
}

func (app *registry) retrieveList(names []string) ([]symbols.Symbol, error) {
	output := []symbols.Symbol{}
	for _, oneName := range names {
		ins, err := app.retrieve(oneName)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return output, nil
}
