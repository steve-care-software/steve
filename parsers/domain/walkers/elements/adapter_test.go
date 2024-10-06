package elements

import (
	"fmt"
	"testing"
)

type testValue struct {
	First  string
	Second string
}

func TestAdapter_Success(t *testing.T) {
	element := Element{
		ElementFn: func(input any) (any, error) {
			return input, nil
		},
		TokenList: &TokenList{
			MapFn: func(elementName string, mp map[string][]any) (any, error) {
				return &testValue{
					First:  string(mp["variableName"][0].(string)),
					Second: string(mp["variableComplex"][0].(string)),
				}, nil
			},
			List: map[string]SelectedTokenList{
				"variableName": {
					SelectorScript: []byte(`
						v1;
						name: mySelector;
						variableName[0][0];
					`),

					Node: &Node{
						Element: &Element{
							ElementFn: func(input any) (any, error) {
								return string(input.([]byte)), nil
							},
						},
					},
				},
				"variableComplex": {
					SelectorScript: []byte(`
						v1;
						name: mySelector;
						variableComplex[0][0];
					`),
					Node: &Node{
						Element: &Element{
							ElementFn: func(input any) (any, error) {
								return string(input.([]byte)), nil
							},
						},
					},
				},
			},
		},
	}

	adapter := NewAdapter()
	retWalker, err := adapter.ToWalker(element)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n%v\n", retWalker)
}
