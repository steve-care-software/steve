package modifications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
	resourcesAdapter    resources.Adapter
	builder             Builder
	modificationBuilder ModificationBuilder
}

func createAdapter(
	resourcesAdapter resources.Adapter,
	builder Builder,
	modificationBuilder ModificationBuilder,
) Adapter {
	out := adapter{
		resourcesAdapter:    resourcesAdapter,
		builder:             builder,
		modificationBuilder: modificationBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *adapter) InstancesToBytes(ins Modifications) ([]byte, error) {
	list := ins.List()
	amount := uint64(len(list))
	output := pointers.Uint64ToBytes(amount)

	for _, oneModification := range list {
		retBytes, err := app.InstanceToBytes(oneModification)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstances converts bytes to instances
func (app *adapter) BytesToInstances(data []byte) (Modifications, []byte, error) {
	pAmount, err := pointers.BytesToUint64(data[0:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	amount := int(*pAmount)
	remaining := data[pointers.Uint64Size:]
	list := []Modification{}
	for i := 0; i < amount; i++ {
		retModification, retRemaining, err := app.BytesToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		remaining = retRemaining
		list = append(list, retModification)
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
func (app *adapter) InstanceToBytes(ins Modification) ([]byte, error) {
	output := []byte{}
	if ins.IsInsert() {
		output = append(output, flagInsert)
		retBytes, err := app.resourcesAdapter.InstanceToBytes(ins.Insert())
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	if ins.IsSave() {
		output = append(output, flagSave)
		retBytes, err := app.resourcesAdapter.InstanceToBytes(ins.Save())
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	if ins.IsDelete() {
		output = append(output, flagDelete)
		strBytes := []byte(ins.Delete())
		amountInBytes := pointers.Uint64ToBytes(uint64(len(strBytes)))
		output = append(output, amountInBytes...)
		output = append(output, strBytes...)
	}

	return output, nil
}

// BytesToInstance converts bytes to instance
func (app *adapter) BytesToInstance(data []byte) (Modification, []byte, error) {
	builder := app.modificationBuilder.Create()
	if len(data) < 1 {
		return nil, nil, errors.New("the data must contain at least 1 byte in order to read the modification flag")
	}

	flag := data[0]
	remaining := data[1:]
	if flag == flagInsert {
		retResource, retRemaining, err := app.resourcesAdapter.BytesToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		builder.WithInsert(retResource)
		remaining = retRemaining
	}

	if flag == flagSave {
		retResource, retRemaining, err := app.resourcesAdapter.BytesToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		builder.WithSave(retResource)
		remaining = retRemaining
	}

	if flag == flagDelete {
		if len(remaining) < pointers.Uint64Size {
			str := fmt.Sprintf("there must be at least %d bytes in order to retrieve the delete's identifier length", pointers.Uint64Size)
			return nil, nil, errors.New(str)
		}

		pLength, err := pointers.BytesToUint64(remaining[0:pointers.Uint64Size])
		if err != nil {
			return nil, nil, err
		}

		length := int(*pLength)
		remaining = remaining[pointers.Uint64Size:]
		if len(remaining) < length {
			str := fmt.Sprintf("there must be at least %d bytes in order retrieve the delete's identifier, %d remaining", length, len(remaining))
			return nil, nil, errors.New(str)
		}

		builder.WithDelete(string(remaining[:length]))
		remaining = remaining[length:]
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}
