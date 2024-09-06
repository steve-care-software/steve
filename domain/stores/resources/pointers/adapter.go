package pointers

import (
	"errors"
	"fmt"
)

type adapter struct {
	builder        Builder
	pointerBuilder PointerBuilder
}

func createAdapter(
	builder Builder,
	pointerBuilder PointerBuilder,
) Adapter {
	out := adapter{
		builder:        builder,
		pointerBuilder: pointerBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *adapter) InstancesToBytes(ins Pointers) ([]byte, error) {
	list := ins.List()
	amount := uint64(len(list))
	output := uint64ToBytes(amount)

	for _, onePointer := range list {
		retBytes, err := app.InstanceToBytes(onePointer)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstances converts bytes to instances
func (app *adapter) BytesToInstances(data []byte) (Pointers, []byte, error) {
	pAmount, err := bytesToUint64(data[0:uint64Size])
	if err != nil {
		return nil, nil, err
	}

	amount := int(*pAmount)
	remaining := data[uint64Size:]
	list := []Pointer{}
	for i := 0; i < amount; i++ {
		retPointer, retRemaining, err := app.BytesToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		remaining = retRemaining
		list = append(list, retPointer)
	}

	ins, err := app.builder.Create().
		WithList(list).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

// InstanceToBytes converts instance to bytes
func (app *adapter) InstanceToBytes(ins Pointer) ([]byte, error) {
	output := uint64ToBytes(uint64(ins.Index()))
	lengthBytes := uint64ToBytes(uint64(ins.Length()))
	return append(output, lengthBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *adapter) BytesToInstance(data []byte) (Pointer, []byte, error) {
	expectation := uint64Size * 2
	if len(data) < expectation {
		str := fmt.Sprintf("there must be at least %d bytes in order to convert them to a Pointer instance", expectation)
		return nil, nil, errors.New(str)
	}

	pIndex, err := bytesToUint64(data[0:uint64Size])
	if err != nil {
		return nil, nil, err
	}

	pLength, err := bytesToUint64(data[uint64Size:expectation])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.pointerBuilder.Create().
		WithIndex(uint(*pIndex)).
		WithLength(uint(*pLength)).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, data[expectation:], nil
}
