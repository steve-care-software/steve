package roots

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
	hashAdapter hash.Adapter
	builder     Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *adapter) ToBytes(ins Root) ([]byte, error) {
	output := pointers.Uint64ToBytes(ins.Amount())
	output = append(output, ins.Owner().Bytes()...)
	return output, nil
}

// ToInstance converts bytes to instance
func (app *adapter) ToInstance(data []byte) (Root, []byte, error) {
	if len(data) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(data))
		return nil, nil, errors.New(str)
	}

	pAmount, err := pointers.BytesToUint64(data[:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	remaining := data[pointers.Uint64Size:]
	if len(remaining) < hash.Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, hash.Size, len(remaining))
		return nil, nil, errors.New(str)
	}

	pOwner, err := app.hashAdapter.FromBytes(remaining[:hash.Size])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.builder.Create().
		WithAmount(*pAmount).
		WithOwner(*pOwner).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining[hash.Size:], nil
}
