package schemas

import (
	"testing"

	"github.com/steve-care-software/steve/parsers/applications"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

func TestGrammar_executeSuites_Success(t *testing.T) {
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

	application := applications.NewApplication()
	err = application.Suites(retGrammar)
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}
}
