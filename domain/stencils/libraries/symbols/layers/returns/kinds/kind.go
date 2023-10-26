package kinds

type kind struct {
	isContinue bool
	isPrompt   bool
	execute    []string
}

func createKindWithContinue() Kind {
	return createKindInternally(true, false, nil)
}

func createKindWithPrompt() Kind {
	return createKindInternally(false, true, nil)
}

func createKindWithExecute(
	execute []string,
) Kind {
	return createKindInternally(false, false, execute)
}

func createKindInternally(
	isContinue bool,
	isPrompt bool,
	execute []string,
) Kind {
	out := kind{
		isContinue: isContinue,
		isPrompt:   isPrompt,
		execute:    execute,
	}

	return &out
}

// IsContinue returns true if continue, false otherwise
func (obj *kind) IsContinue() bool {
	return obj.isContinue
}

// IsPrompt returns true if prompt, false otherwise
func (obj *kind) IsPrompt() bool {
	return obj.isPrompt
}

// IsExecute returns true if execute, false otherwise
func (obj *kind) IsExecute() bool {
	return obj.execute != nil
}

// Execute returns the execute commands, if any
func (obj *kind) Execute() []string {
	return obj.execute
}
