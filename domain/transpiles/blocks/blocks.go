package blocks

type blocks struct {
	list []Block
}

func createBlocks(
	list []Block,
) Blocks {
	out := blocks{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *blocks) List() []Block {
	return obj.list
}
