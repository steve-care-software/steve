package commits

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
	modificationsAdapter modifications.Adapter
	hashAdapter          hash.Adapter
	builder              Builder
	commitBuilder        CommitBuilder
}

func createAdapter(
	modificationsAdapter modifications.Adapter,
	hashAdapter hash.Adapter,
	builder Builder,
	commitBuilder CommitBuilder,
) Adapter {
	out := adapter{
		modificationsAdapter: modificationsAdapter,
		hashAdapter:          hashAdapter,
		builder:              builder,
		commitBuilder:        commitBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *adapter) InstancesToBytes(ins Commits) ([]byte, error) {
	list := ins.List()
	amount := uint64(len(list))
	output := pointers.Uint64ToBytes(amount)

	for _, oneCommit := range list {
		retBytes, err := app.InstanceToBytes(oneCommit)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstances converts bytes to instances
func (app *adapter) BytesToInstances(data []byte) (Commits, []byte, error) {
	pAmount, err := pointers.BytesToUint64(data[0:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	amount := int(*pAmount)
	remaining := data[pointers.Uint64Size:]
	list := []Commit{}
	for i := 0; i < amount; i++ {
		retCommit, retRemaining, err := app.BytesToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		remaining = retRemaining
		list = append(list, retCommit)
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
func (app *adapter) InstanceToBytes(ins Commit) ([]byte, error) {
	modifications := ins.Modifications()
	retModificationBytes, err := app.modificationsAdapter.InstancesToBytes(modifications)
	if err != nil {
		return nil, err
	}

	hasParentByte := byte(0)
	parentBytes := []byte{}
	if ins.HasParent() {
		hasParentByte = byte(1)
		parentBytes = ins.Parent().Bytes()
	}

	output := retModificationBytes
	output = append(output, hasParentByte)
	output = append(output, parentBytes...)

	return output, nil
}

// BytesToInstance converts bytes to instance
func (app *adapter) BytesToInstance(data []byte) (Commit, []byte, error) {
	retModifications, retRemaining, err := app.modificationsAdapter.BytesToInstances(data)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) < 1 {
		return nil, nil, errors.New("the data was expected to contain 1 byte that represents the flag that tells if there is a parent hash")
	}

	// true
	remaining := retRemaining[1:]
	builder := app.commitBuilder.Create().WithModifications(retModifications)
	if retRemaining[0] == 1 {
		if len(remaining) < hash.Size {
			str := fmt.Sprintf("the data was expected to contain at least %d bytes to form the parent hash, %d provided", hash.Size, len(remaining))
			return nil, nil, errors.New(str)
		}

		pHash, err := app.hashAdapter.FromBytes(remaining[:hash.Size])
		if err != nil {
			return nil, nil, err
		}

		remaining = remaining[hash.Size:]
		builder.WithParent(*pHash)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}
