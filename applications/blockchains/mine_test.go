package blockchains

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/hash"
)

func TestMine_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pFlag, err := hashAdapter.FromBytes([]byte("flag"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pScript, err := hashAdapter.FromBytes([]byte("script"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pOtherScript, err := hashAdapter.FromBytes([]byte("other script"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pubKey, pk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	firstEntry := entries.NewEntryForTests(*pFlag, *pScript, 22)
	firstSignature := ed25519.Sign(pk, firstEntry.Hash().Bytes())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	secondEntry := entries.NewEntryForTests(*pFlag, *pOtherScript, 34)
	secondSignature := ed25519.Sign(pk, secondEntry.Hash().Bytes())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	trx := transactions.NewTransactionsForTests([]transactions.Transaction{
		transactions.NewTransactionForTests(
			firstEntry,
			firstSignature,
			pubKey,
		),
		transactions.NewTransactionForTests(
			secondEntry,
			secondSignature,
			pubKey,
		),
	})

	requestedDifficulty := uint8(3)
	miningValue := uint8(0)
	result, err := mine(hashAdapter, trx, requestedDifficulty, miningValue)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pHash, err := hashAdapter.FromMultiBytes([][]byte{
		trx.Hash().Bytes(),
		result,
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	prefix := []byte{
		miningValue,
		miningValue,
		miningValue,
	}

	if !bytes.HasPrefix(pHash.Bytes(), prefix) {
		t.Errorf("the returned result is invalid")
		return
	}

}
