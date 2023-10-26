package layers

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/results/executions/inits"
)

type computer struct {
}

func (app *computer) init(init inits.Init) (*uint, error) {
	return nil, nil
}

func (app *computer) retrieve(context uint, name string) ([]byte, error) {
	return nil, nil
}

func (app *computer) retrieveLayer(context uint, name string) (layers.Layer, error) {
	return nil, nil
}

func (app *computer) reduce(context uint, name string, length uint8) ([]byte, error) {
	return nil, nil
}

func (app *computer) assign(context uint, name string, values []byte) error {
	return nil
}
