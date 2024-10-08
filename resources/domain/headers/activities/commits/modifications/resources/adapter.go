package resources

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
	pointerAdapter  pointers.Adapter
	builder         Builder
	resourceBuilder ResourceBuilder
}

func createAdapter(
	pointerAdapter pointers.Adapter,
	builder Builder,
	resourceBuilder ResourceBuilder,
) Adapter {
	out := adapter{
		pointerAdapter:  pointerAdapter,
		builder:         builder,
		resourceBuilder: resourceBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *adapter) InstancesToBytes(ins Resources) ([]byte, error) {
	list := ins.List()
	amount := uint64(len(list))
	output := pointers.Uint64ToBytes(amount)

	for _, oneResource := range list {
		retBytes, err := app.InstanceToBytes(oneResource)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstances converts bytes to instances
func (app *adapter) BytesToInstances(data []byte) (Resources, []byte, error) {
	pAmount, err := pointers.BytesToUint64(data[0:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	amount := int(*pAmount)
	remaining := data[pointers.Uint64Size:]
	list := []Resource{}
	for i := 0; i < amount; i++ {
		retResource, retRemaining, err := app.BytesToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		remaining = retRemaining
		list = append(list, retResource)
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
func (app *adapter) InstanceToBytes(ins Resource) ([]byte, error) {
	identifierBytes := []byte(ins.Identifier())
	length := len(identifierBytes)
	output := pointers.Uint64ToBytes(uint64(length))
	output = append(output, identifierBytes...)

	pointers := ins.Pointer()
	retBytes, err := app.pointerAdapter.ToBytes(pointers)
	if err != nil {
		return nil, err
	}

	return append(output, retBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *adapter) BytesToInstance(data []byte) (Resource, []byte, error) {
	if len(data) < pointers.Uint64Size {
		str := fmt.Sprintf("there must be at least %d bytes in order to retrieve the identifier's length", pointers.Uint64Size)
		return nil, nil, errors.New(str)
	}

	pLength, err := pointers.BytesToUint64(data[0:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	length := int(*pLength)
	remaining := data[pointers.Uint64Size:]
	if len(remaining) < length {
		str := fmt.Sprintf("there must be at least %d bytes in order retrieve the identifier, %d remaining", length, len(remaining))
		return nil, nil, errors.New(str)
	}

	identifier := string(remaining[0:length])
	retPointer, retRemaining, err := app.pointerAdapter.ToInstance(remaining[length:])
	if err != nil {
		return nil, nil, err
	}

	resource, err := app.resourceBuilder.Create().
		WithIdentifier(identifier).
		WithPointer(retPointer).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return resource, retRemaining, nil
}
