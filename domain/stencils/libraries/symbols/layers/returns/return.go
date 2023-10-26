package returns

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"

type ret struct {
	output []byte
	kind   kinds.Kind
}

func createReturn(
	output []byte,
	kind kinds.Kind,
) Return {
	out := ret{
		output: output,
		kind:   kind,
	}

	return &out
}

// Output returns the output
func (obj *ret) Output() []byte {
	return obj.output
}

// Kind returns the kind
func (obj *ret) Kind() kinds.Kind {
	return obj.kind
}
