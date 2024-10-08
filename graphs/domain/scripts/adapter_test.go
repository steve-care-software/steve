package scripts

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/parsers/applications"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

func TestAdapter_Success(t *testing.T) {
	grammarInput := fetchGrammarInput()
	grammarParserAdapter := grammars.NewAdapter()
	retGrammar, retRemaining, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) > 0 {
		t.Errorf("the remaining script was expected to be empty, this was returned: %s", retRemaining)
		return
	}

	application, _ := applications.NewBuilder().Create().Now()
	err = application.Suites(retGrammar)
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}

	adapter, err := NewAdapterBuilder().Create().WithGramar(retGrammar).Now()
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}

	input := []byte(`
		head:
			engine: v1;
			name: mySchema;
			access: 
				read: .first .second (0.2);
				write: 
					.first .again;
					review: .first .second .third (0.1);
				;
			;
		;

		son;
		father;
		grandFather;
		grandGrandFather;

		father[0,3](son[1,]): .son .father
						| .father .grandFather
						| .grandFather .grandGrandFather
						---
							mySuite[ .mySchema[son] .grandGrandFather]:
								!(.son .father .grandFather .grandGrandFather);
								(.son .father .grandFather .grandGrandFather);
							;
						;

		grandFather(grandSon[2,]): .son .grandFather
								| .father .grandGrandFather
								;
	`)

	retScript, retRemaining, err := adapter.ToScript(input)
	if err != nil {
		t.Errorf("there was an error while converting the input to a Script instance: %s", err.Error())
		return
	}

	if len(retRemaining) > 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	expectedPoints := []string{
		"son",
		"father",
		"grandFather",
		"grandGrandFather",
	}

	if !reflect.DeepEqual(expectedPoints, retScript.Schema().Points()) {
		t.Errorf("the returned points are invalid")
		return
	}
}
