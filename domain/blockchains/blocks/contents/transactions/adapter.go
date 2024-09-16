package transactions

import (
	"crypto/ed25519"
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
	entryAdapter       entries.Adapter
	builder            Builder
	transactionBuilder TransactionBuilder
}

func createAdapter(
	entryAdapter entries.Adapter,
	builder Builder,
	transactionBuilder TransactionBuilder,
) Adapter {
	out := adapter{
		entryAdapter:       entryAdapter,
		builder:            builder,
		transactionBuilder: transactionBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *adapter) InstancesToBytes(ins Transactions) ([]byte, error) {
	list := ins.List()
	amount := uint64(len(list))
	output := pointers.Uint64ToBytes(amount)

	for _, oneTrx := range list {
		retBytes, err := app.InstanceToBytes(oneTrx)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstances converts bytes to instances
func (app *adapter) BytesToInstances(data []byte) (Transactions, []byte, error) {
	if len(data) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(data))
		return nil, nil, errors.New(str)
	}

	pAmount, err := pointers.BytesToUint64(data[:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	amount := int(*pAmount)
	remaining := data[pointers.Uint64Size:]
	list := []Transaction{}
	for i := 0; i < amount; i++ {
		retTrx, retRemaining, err := app.BytesToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		remaining = retRemaining
		list = append(list, retTrx)
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
func (app *adapter) InstanceToBytes(ins Transaction) ([]byte, error) {
	entryBytes, err := app.entryAdapter.ToBytes(ins.Entry())
	if err != nil {
		return nil, err
	}

	signatureBytes := ins.Signature()
	pubKeyBytes := ins.PublicKey()

	output := entryBytes
	output = append(output, signatureBytes...)
	output = append(output, pubKeyBytes...)
	return output, nil
}

// BytesToInstance converts bytes to instance
func (app *adapter) BytesToInstance(data []byte) (Transaction, []byte, error) {
	retEntry, retRemaining, err := app.entryAdapter.ToInstance(data)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) < ed25519.SignatureSize {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, ed25519.SignatureSize, len(retRemaining))
		return nil, nil, errors.New(str)
	}

	signatureBytes := retRemaining[:ed25519.SignatureSize]
	remaining := retRemaining[ed25519.SignatureSize:]
	if len(remaining) < ed25519.PublicKeySize {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, ed25519.PublicKeySize, len(remaining))
		return nil, nil, errors.New(str)
	}

	pubKey := remaining[:ed25519.PublicKeySize]
	ins, err := app.transactionBuilder.Create().
		WithEntry(retEntry).
		WithSignature(signatureBytes).
		WithPublicKey(pubKey).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, remaining[ed25519.PublicKeySize:], nil
}
