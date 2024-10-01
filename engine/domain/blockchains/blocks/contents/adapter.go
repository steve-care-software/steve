package contents

import (
	"crypto/ed25519"
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents/transactions"
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
	output = append(output, ins.Parent().Bytes()...)
	output = append(output, ins.Miner()...)
	output = append(output, ins.Commit().Bytes()...)
	return output, nil
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

	remaining := retRemaining[hash.Size:]
	if len(remaining) < hash.Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, hash.Size, len(remaining))
		return nil, nil, errors.New(str)
	}

	miner := remaining[:ed25519.PublicKeySize]
	remaining = remaining[ed25519.PublicKeySize:]
	if len(remaining) < hash.Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, hash.Size, len(remaining))
		return nil, nil, errors.New(str)
	}

	pCommit, err := app.hashAdapter.FromBytes(remaining[:hash.Size])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.builder.Create().
		WithTransactions(retTrx).
		WithParent(*pHash).
		WithMiner(miner).
		WithCommit(*pCommit).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining[hash.Size:], nil
}
