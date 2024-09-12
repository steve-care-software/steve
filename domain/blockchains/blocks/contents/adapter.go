package contents

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

type adapter struct {
	trxAdapter  transactions.Adapter
	hashAdapter hash.Adapter
	builder     Builder
}

func createAdapter(
	trxAdapter transactions.Adapter,
	hashAdapter hash.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		trxAdapter:  trxAdapter,
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *adapter) ToBytes(ins Content) ([]byte, error) {
	trxBytes, err := app.trxAdapter.InstancesToBytes(ins.Transactions())
	if err != nil {
		return nil, err
	}
	output := trxBytes
	return append(output, ins.Parent().Bytes()...), nil
}

// ToInstance converts bytes to instance
func (app *adapter) ToInstance(data []byte) (Content, []byte, error) {
	retTrx, retRemaining, err := app.trxAdapter.BytesToInstances(data)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) < hash.Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, hash.Size, len(retRemaining))
		return nil, nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(retRemaining[:hash.Size])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.builder.Create().
		WithTransactions(retTrx).
		WithParent(*pHash).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining[hash.Size:], nil
}
