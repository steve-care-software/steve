package instructions

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

type tokens struct {
	list []Token
	mp   map[string][]Token
}

func createTokens(
	list []Token,
	mp map[string][]Token,
) Tokens {
	out := tokens{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of token
func (obj *tokens) List() []Token {
	return obj.list
}

// Fetch fetches a token by name and index
func (obj *tokens) Fetch(name string, idx uint) (Token, error) {
	if ins, ok := obj.mp[name]; ok {
		length := uint(len(ins))
		if idx >= length {
			str := fmt.Sprintf("the token (%s) could not be found at index (%d), its length is: %d", name, idx, length)
			return nil, errors.New(str)
		}

		return ins[idx], nil
	}

	str := fmt.Sprintf("the token (name: %s) does not exists", name)
	return nil, errors.New(str)
}

// IsBalanceValid validates the tokens against the balance
func (obj *tokens) IsBalanceValid(balance balances.Balance) bool {
	list := balance.Lines()
	for _, oneSelectors := range list {
		operationIsValid := true
		selectorsList := oneSelectors.List()
		for _, oneSelector := range selectorsList {
			isValid := obj.IsSelectorValid(oneSelector)
			if !isValid {
				operationIsValid = false
			}
		}

		if operationIsValid {
			return true
		}
	}

	return false
}

// IsSelectorValid validates the tokens against the selector
func (obj *tokens) IsSelectorValid(selector selectors.Selector) bool {
	chain := selector.Chain()
	isChainValid := obj.IsChainValid(chain)
	if selector.IsNot() {
		return !isChainValid
	}

	return isChainValid
}

// IsChainValid validates the tokens against the chain
func (obj *tokens) IsChainValid(chain chains.Chain) bool {
	name := chain.Element().Name()
	if chain.HasToken() {
		token := chain.Token()
		tokenIndex := token.Index()
		retASTToken, err := obj.Fetch(name, tokenIndex)
		if err != nil {
			return false
		}

		if token.HasElement() {
			element := token.Element()
			elementIndex := element.Index()
			retASTElement, err := retASTToken.Elements().Fetch(elementIndex)
			if err != nil {
				return false
			}

			if element.HasChain() {
				chain := element.Chain()
				return retASTElement.IsChainValid(chain)
			}

			return true
		}

		return true
	}

	_, err := obj.Fetch(name, 0)
	return err == nil
}
