package contents

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/databases/blockchains/domain/blocks/contents/transactions"
	"github.com/steve-care-software/steve/databases/blockchains/domain/blocks/contents/transactions/entries"
)

func TestAdapter_withRemaining_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pParent, err := hashAdapter.FromBytes([]byte("parent"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

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

	pCommit, err := hashAdapter.FromBytes([]byte("commit hash"))
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

	trx := NewContentForTests(
		transactions.NewTransactionsForTests([]transactions.Transaction{
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
		}),
		*pParent,
		pubKey,
		*pCommit,
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(trx)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining data")
	retTrx, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(trx, retTrx) {
		t.Errorf("the returned transaction is invalid")
		return
	}
}
