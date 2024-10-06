package elements

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/parsers/domain/queries"
	"github.com/steve-care-software/steve/parsers/domain/walkers"
)

type adapter struct {
	queryAdapter             queries.Adapter
	builder                  walkers.Builder
	tokenListBuilder         walkers.TokenListBuilder
	selectedTokenListBuilder walkers.SelectedTokenListBuilder
	tokenBuilder             walkers.TokenBuilder
	nodeBuilder              walkers.NodeBuilder
}

func createAdapter(
	queryAdapter queries.Adapter,
	builder walkers.Builder,
	tokenListBuilder walkers.TokenListBuilder,
	selectedTokenListBuilder walkers.SelectedTokenListBuilder,
	tokenBuilder walkers.TokenBuilder,
	nodeBuilder walkers.NodeBuilder,
) Adapter {
	out := adapter{
		queryAdapter:             queryAdapter,
		builder:                  builder,
		tokenListBuilder:         tokenListBuilder,
		selectedTokenListBuilder: selectedTokenListBuilder,
		tokenBuilder:             tokenBuilder,
		nodeBuilder:              nodeBuilder,
	}

	return &out
}

// ToWalker converts an element to a walker instance
func (app *adapter) ToWalker(ins Element) (walkers.Walker, error) {
	builder := app.builder.Create().WithFn(ins.ElementFn)
	if ins.TokenList != nil {
		retTokenList, err := app.tokenList(*ins.TokenList)
		if err != nil {
			return nil, err
		}

		builder.WithList(retTokenList)
	}

	return builder.Now()
}

func (app *adapter) tokenList(ins TokenList) (walkers.TokenList, error) {
	list := []walkers.SelectedTokenList{}
	for name, oneSelectedTokenList := range ins.List {
		retSelectedTokenList, err := app.selectedTokenList(name, oneSelectedTokenList)
		if err != nil {
			return nil, err
		}

		list = append(list, retSelectedTokenList)
	}
	return app.tokenListBuilder.Create().
		WithFn(ins.MapFn).
		WithList(list).
		Now()
}

func (app *adapter) selectedTokenList(name string, ins SelectedTokenList) (walkers.SelectedTokenList, error) {
	query, remaining, err := app.queryAdapter.ToQuery(ins.SelectorScript)
	if err != nil {
		return nil, err
	}

	if len(remaining) > 0 {
		str := fmt.Sprintf("the script (%s) contains a remaining (%s)", ins.SelectorScript, remaining)
		return nil, errors.New(str)
	}

	chain := query.Chain()
	builder := app.selectedTokenListBuilder.Create().
		WithChain(chain).
		WithName(name)

	if ins.Node != nil {
		retNode, err := app.node(*ins.Node)
		if err != nil {
			return nil, err
		}

		builder.WithNode(retNode)
	}

	return builder.Now()
}

func (app *adapter) token(ins Token) (walkers.Token, error) {
	builder := app.tokenBuilder.Create().WithFn(ins.ListFn)
	if ins.Next != nil {
		retWalker, err := app.ToWalker(*ins.Next)
		if err != nil {
			return nil, err
		}

		builder.WithNext(retWalker)
	}

	return builder.Now()
}

func (app *adapter) node(ins Node) (walkers.Node, error) {
	builder := app.nodeBuilder.Create()
	if ins.Element != nil {
		retWalker, err := app.ToWalker(*ins.Element)
		if err != nil {
			return nil, err
		}

		builder.WithElement(retWalker)
	}

	if ins.Token != nil {
		retToken, err := app.token(*ins.Token)
		if err != nil {
			return nil, err
		}

		builder.WithToken(retToken)
	}

	if ins.TokenList != nil {
		retTokenList, err := app.tokenList(*ins.TokenList)
		if err != nil {
			return nil, err
		}

		builder.WithTokenList(retTokenList)
	}

	return builder.Now()
}
