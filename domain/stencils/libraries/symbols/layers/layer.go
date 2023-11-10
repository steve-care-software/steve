package layers

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters"
	return_expectations "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/expectations"
)

type layer struct {
	hash       hash.Hash
	input      string
	executions Executions
	ret        return_expectations.Expectation
	params     parameters.Parameters
	suites     Suites
}

func createLayer(
	hash hash.Hash,
	input string,
	executions Executions,
	ret return_expectations.Expectation,
) Layer {
	return createLayerInternally(hash, input, executions, ret, nil, nil)
}

func createLayerWithParams(
	hash hash.Hash,
	input string,
	executions Executions,
	ret return_expectations.Expectation,
	params parameters.Parameters,
) Layer {
	return createLayerInternally(hash, input, executions, ret, params, nil)
}

func createLayerWithSuites(
	hash hash.Hash,
	input string,
	executions Executions,
	ret return_expectations.Expectation,
	suites Suites,
) Layer {
	return createLayerInternally(hash, input, executions, ret, nil, suites)
}

func createLayerWithParamsAndSuites(
	hash hash.Hash,
	input string,
	executions Executions,
	ret return_expectations.Expectation,
	params parameters.Parameters,
	suites Suites,
) Layer {
	return createLayerInternally(hash, input, executions, ret, params, suites)
}

func createLayerInternally(
	hash hash.Hash,
	input string,
	executions Executions,
	ret return_expectations.Expectation,
	params parameters.Parameters,
	suites Suites,
) Layer {
	out := layer{
		hash:       hash,
		input:      input,
		executions: executions,
		ret:        ret,
		params:     params,
		suites:     suites,
	}

	return &out
}

// Hash returns the hash
func (obj *layer) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *layer) Input() string {
	return obj.input
}

// Executions returns the executions
func (obj *layer) Executions() Executions {
	return obj.executions
}

// Return returns the return
func (obj *layer) Return() return_expectations.Expectation {
	return obj.ret
}

// HasParams returns true if there is params, false otherwise
func (obj *layer) HasParams() bool {
	return obj.params != nil
}

// Params returns the params, if any
func (obj *layer) Params() parameters.Parameters {
	return obj.params
}

// HasSuites returns true if there is suites, false otherwise
func (obj *layer) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *layer) Suites() Suites {
	return obj.suites
}
