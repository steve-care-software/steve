package scripts

import (
	"testing"

	"github.com/steve-care-software/steve/applications/chains/interpreters"
	"github.com/steve-care-software/steve/domain/programs/grammars"
)

func TestParserAdapter_execTestSuites_Success(t *testing.T) {
	grammarInput := FetchGrammarInput()
	grammarParserAdapter := grammars.NewParserAdapter()
	retGrammar, _, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	application := interpreters.NewApplication()
	err = application.Suites(retGrammar)
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}
}
