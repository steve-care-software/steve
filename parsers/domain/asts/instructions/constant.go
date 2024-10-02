package instructions

import "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"

type constant struct {
	name  string
	value []byte
}

func createConstant(
	name string,
	value []byte,
) Constant {
	out := constant{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *constant) Name() string {
	return obj.name
}

// Value returns the value
func (obj *constant) Value() []byte {
	return obj.value
}

// IsChainValid validates the constant against the chain
func (obj *constant) IsChainValid(chain chains.Chain) bool {
	name := chain.Element().Name()
	if chain.HasToken() {
		return false
	}

	return obj.name == name
}
