package tokens

import "github.com/steve-care-software/steve/parsers/domain/grammars/constants/tokens/elements"

type token struct {
	element elements.Element
	amount  uint
}

func createToken(
	element elements.Element,
	amount uint,
) Token {
	out := token{
		element: element,
		amount:  amount,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	if obj.element.IsConstant() {
		return obj.element.Constant()
	}

	return obj.element.Rule()
}

// Element returns the element
func (obj *token) Element() elements.Element {
	return obj.element
}

// Amount returns the amount
func (obj *token) Amount() uint {
	return obj.amount
}
