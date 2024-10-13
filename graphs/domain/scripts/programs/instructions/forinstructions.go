package instructions

type forInstructions struct {
	list []ForInstruction
}

func createForInstructions(
	list []ForInstruction,
) ForInstructions {
	out := forInstructions{
		list: list,
	}

	return &out
}

// List returns the list of forInstruction
func (obj *forInstructions) List() []ForInstruction {
	return obj.list
}
