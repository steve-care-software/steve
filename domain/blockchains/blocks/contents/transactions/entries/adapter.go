package entries

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
func (app *adapter) ToBytes(ins Entry) ([]byte, error) {
	output := ins.Flag().Bytes()
	output = append(output, ins.Script().Bytes()...)
	output = append(output, pointers.Uint64ToBytes(ins.Fees())...)
	return output, nil
}

// ToInstance converts bytes to instance
func (app *adapter) ToInstance(data []byte) (Entry, error) {
	if len(data) < hash.Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, hash.Size, len(data))
		return nil, errors.New(str)
	}

	pFlag, err := app.hashAdapter.FromBytes(data[:hash.Size])
	if err != nil {
		return nil, err
	}

	remaining := data[hash.Size:]
	if len(data) < hash.Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, hash.Size, len(remaining))
		return nil, errors.New(str)
	}

	pScript, err := app.hashAdapter.FromBytes(remaining[:hash.Size])
	if err != nil {
		return nil, err
	}

	remaining = remaining[hash.Size:]
	if len(data) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(remaining))
		return nil, errors.New(str)
	}

	pFees, err := pointers.BytesToUint64(remaining[:pointers.Uint64Size])
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithFlag(*pFlag).
		WithScript(*pScript).
		WithFees(*pFees).
		Now()
}
