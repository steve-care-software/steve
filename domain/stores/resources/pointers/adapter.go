package pointers

import (
	"errors"
	"fmt"
)

type adapter struct {
	builder Builder
}

func createAdapter(
	builder Builder,
) Adapter {
	out := adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *adapter) ToBytes(ins Pointer) ([]byte, error) {
	output := Uint64ToBytes(uint64(ins.Index()))
	lengthBytes := Uint64ToBytes(uint64(ins.Length()))
	return append(output, lengthBytes...), nil
}

// ToInstance converts bytes to instance
func (app *adapter) ToInstance(data []byte) (Pointer, []byte, error) {
	expectation := Uint64Size * 2
	if len(data) < expectation {
		str := fmt.Sprintf("there must be at least %d bytes in order to convert them to a Pointer instance", expectation)
		return nil, nil, errors.New(str)
	}

	pIndex, err := BytesToUint64(data[0:Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	pLength, err := BytesToUint64(data[Uint64Size:expectation])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.builder.Create().
		WithIndex(uint(*pIndex)).
		WithLength(uint(*pLength)).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, data[expectation:], nil
}
