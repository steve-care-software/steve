package blocks

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
	contentAdapter contents.Adapter
	builder        Builder
	blockBuilder   BlockBuilder
}

func createAdapter(
	contentAdapter contents.Adapter,
	builder Builder,
	blockBuilder BlockBuilder,
) Adapter {
	out := adapter{
		contentAdapter: contentAdapter,
		builder:        builder,
		blockBuilder:   blockBuilder,
	}

	return &out
}

// InstancesToBytes converts blocks to bytes
func (app *adapter) InstancesToBytes(ins Blocks) ([]byte, error) {
	list := ins.List()
	amount := uint64(len(list))
	output := pointers.Uint64ToBytes(amount)

	for _, oneBlock := range list {
		retBytes, err := app.InstanceToBytes(oneBlock)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstances converts bytes to instances
func (app *adapter) BytesToInstances(data []byte) (Blocks, []byte, error) {
	pAmount, err := pointers.BytesToUint64(data[0:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	amount := int(*pAmount)
	remaining := data[pointers.Uint64Size:]
	list := []Block{}
	for i := 0; i < amount; i++ {
		retBlock, retRemaining, err := app.BytesToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		remaining = retRemaining
		list = append(list, retBlock)
	}

	ins, err := app.builder.Create().
		WithList(list).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

// InstanceToBytes converts block to bytes
func (app *adapter) InstanceToBytes(ins Block) ([]byte, error) {
	contentBytes, err := app.contentAdapter.ToBytes(ins.Content())
	if err != nil {
		return nil, err
	}

	resultBytes := ins.Result()
	lengthBytes := pointers.Uint64ToBytes(uint64(len(resultBytes)))

	output := contentBytes
	output = append(output, lengthBytes...)
	return append(output, resultBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *adapter) BytesToInstance(data []byte) (Block, []byte, error) {
	retContent, retRemaining, err := app.contentAdapter.ToInstance(data)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(retRemaining))
		return nil, nil, errors.New(str)
	}

	pLength, err := pointers.BytesToUint64(retRemaining[:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	remaining := retRemaining[pointers.Uint64Size:]
	if uint64(len(remaining)) < *pLength {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, *pLength, len(remaining))
		return nil, nil, errors.New(str)
	}

	resultBytes := remaining[:*pLength]
	ins, err := app.blockBuilder.Create().
		WithContent(retContent).
		WithResult(resultBytes).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining[*pLength:], nil
}
