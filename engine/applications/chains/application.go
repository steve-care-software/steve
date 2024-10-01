package chains

import (
	transpiles_applications "github.com/steve-care-software/steve/engine/applications/chains/transpiles"
	"github.com/steve-care-software/steve/engine/domain/chains"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles"
	interpreters "github.com/steve-care-software/steve/parsers/applications"
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/programs"
)

type application struct {
	interpreterApp       interpreters.Application
	transpileApp         transpiles_applications.Application
	grammarNFTAdapter    grammars.NFTAdapter
	programParserAdapter asts.ParserAdapter
	transpileNFTAdapter  transpiles.NFTAdapter
}

func createApplication(
	interpreterApp interpreters.Application,
	transpileApp transpiles_applications.Application,
	grammarNFTAdapter grammars.NFTAdapter,
	programParserAdapter asts.ParserAdapter,
	transpileNFTAdapter transpiles.NFTAdapter,
) Application {
	out := application{
		interpreterApp:       interpreterApp,
		transpileApp:         transpileApp,
		grammarNFTAdapter:    grammarNFTAdapter,
		programParserAdapter: programParserAdapter,
		transpileNFTAdapter:  transpileNFTAdapter,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(chain chains.Chain, input []byte) ([]byte, error) {
	nft := chain.Grammar()
	grammar, err := app.grammarNFTAdapter.ToInstance(nft)
	if err != nil {
		return nil, err
	}

	program, _, err := app.programParserAdapter.ToProgram(grammar, input)
	if err != nil {
		return nil, err
	}

	action := chain.Action()
	return app.action(action, program)
}

func (app *application) action(action chains.Action, program programs.Program) ([]byte, error) {
	if action.IsInterpret() {
		interpret := action.Interpret()
		retStack, err := app.interpreterApp.Execute(program)
		if err != nil {
			return nil, err
		}

		variableName := interpret.Variable()
		retBytes, err := retStack.FetchBytes(variableName)
		if err != nil {
			return nil, err
		}

		if interpret.HasNext() {
			next := interpret.Next()
			return app.Execute(next, retBytes)
		}

		return retBytes, nil
	}

	actionTranspile := action.Transpile()
	transpileNFT := actionTranspile.Bridge()
	transpile, err := app.transpileNFTAdapter.ToInstance(transpileNFT)
	if err != nil {
		return nil, err
	}

	grammarNFT := actionTranspile.Target()
	grammar, err := app.grammarNFTAdapter.ToInstance(grammarNFT)
	if err != nil {
		return nil, err
	}

	retBytes, err := app.transpileApp.Execute(program, grammar, transpile)
	if err != nil {
		return nil, err
	}

	if actionTranspile.HasNext() {
		next := actionTranspile.Next()
		return app.Execute(next, retBytes)
	}

	return retBytes, nil
}
