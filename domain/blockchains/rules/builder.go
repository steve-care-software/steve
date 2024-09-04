package rules

import "errors"

type builder struct {
	pMiningValue   *uint8
	baseDifficulty uint8
	incrDiffPerTrx float64
}

func createBuilder() Builder {
	out := builder{
		pMiningValue:   nil,
		baseDifficulty: 0,
		incrDiffPerTrx: 0.0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMiningValue adds a miningValue to the builder
func (app *builder) WithMiningValue(miningValue uint8) Builder {
	app.pMiningValue = &miningValue
	return app
}

// WithBaseDifficulty adds a baseDifficulty to the builder
func (app *builder) WithBaseDifficulty(baseDifficulty uint8) Builder {
	app.baseDifficulty = baseDifficulty
	return app
}

// WithIncreaseDifficultyPerTrx adds an increase difficulty per trx
func (app *builder) WithIncreaseDifficultyPerTrx(incrDiffPerTrx float64) Builder {
	app.incrDiffPerTrx = incrDiffPerTrx
	return app
}

// Now builds a new Rules instance
func (app *builder) Now() (Rules, error) {
	if app.pMiningValue != nil {
		return nil, errors.New("the mining value is mandatory in order to build a Rules instance")
	}

	if app.baseDifficulty <= 0 {
		return nil, errors.New("the base difficulty is mandatory in order to build a Rules instance")
	}

	if app.incrDiffPerTrx <= 0 {
		return nil, errors.New("the increase difficulty per transaction is mandatory in order to build a Rules instance")
	}

	return createRules(
		*app.pMiningValue,
		app.baseDifficulty,
		app.incrDiffPerTrx,
	), nil
}
