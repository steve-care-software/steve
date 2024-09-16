package blocks

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/hash"
)

func TestAdapter_single_withRemaining_Success(t *testing.T) {
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

	block := NewBlockForTests(
		contents.NewContentForTests(
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
		),
		[]byte("this is some result"),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(block)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining data")
	retBlock, retRemaining, err := adapter.BytesToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(block, retBlock) {
		t.Errorf("the returned block is invalid")
		return
	}
}

func TestAdapter_multiple_withRemaining_Success(t *testing.T) {
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

	blocks := NewBlocksForTests([]Block{
		NewBlockForTests(
			contents.NewContentForTests(
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
			),
			[]byte("this is some result"),
		),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(blocks)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining data")
	retBlock, retRemaining, err := adapter.BytesToInstances(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(blocks, retBlock) {
		t.Errorf("the returned blocsk is invalid")
		return
	}
}
