package scripts

// FetchGrammarInput fetches the grammar input
func FetchGrammarInput() []byte {
	return grammarInput()
}

// ToTranspile converts an input to a script instance
type ParserAdapter interface {
	ToTransfer(input []byte) (Script, []byte, error)
}

// Script represents a script
type Script interface {
}
