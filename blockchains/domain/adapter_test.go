package blockchains

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/blockchains/domain/blocks"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents/transactions"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/blockchains/domain/roots"
	"github.com/steve-care-software/steve/blockchains/domain/rules"
	"github.com/steve-care-software/steve/hash"
)

func TestAdapter_withoutDescription_withoutHead_withRemaining_Success(t *testing.T) {
	identifier, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashAdapter := hash.NewAdapter()
	pCommit, err := hashAdapter.FromBytes([]byte("commit"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pubKey, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	blockchain := NewBlockchainForTests(
		identifier,
		"myBlockchain",
		"",
		rules.NewRulesForTests(0, 2, 0.01),
		roots.NewRootForTests(456, pubKey, *pCommit),
		time.Now().UTC(),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(blockchain)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some additional bytes")
	retBlockchain, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytesAgain, err := adapter.ToBytes(retBlockchain)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !bytes.Equal(retBytes, retBytesAgain) {
		t.Errorf("the returned blockchain bytes are invalid")
		return
	}

}

func TestAdapter_withDescription_withoutHead_withRemaining_Success(t *testing.T) {
	identifier, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashAdapter := hash.NewAdapter()
	pCommit, err := hashAdapter.FromBytes([]byte("commit"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pubKey, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	blockchain := NewBlockchainForTests(
		identifier,
		"myBlockchain",
		"this is a description",
		rules.NewRulesForTests(0, 2, 0.01),
		roots.NewRootForTests(456, pubKey, *pCommit),
		time.Now().UTC(),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(blockchain)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some additional bytes")
	retBlockchain, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytesAgain, err := adapter.ToBytes(retBlockchain)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !bytes.Equal(retBytes, retBytesAgain) {
		t.Errorf("the returned blockchain bytes are invalid")
		return
	}

}

func TestAdapter_withDescription_withHead_withRemaining_Success(t *testing.T) {
	identifier, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

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

	pCommit, err := hashAdapter.FromBytes([]byte("commit"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pContentCommmit, err := hashAdapter.FromBytes([]byte("content commit hash"))
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

	blockchain := NewBlockchainWithHeadForTests(
		identifier,
		"myBlockchain",
		"this is a description",
		rules.NewRulesForTests(0, 2, 0.01),
		roots.NewRootForTests(456, pubKey, *pCommit),
		time.Now().UTC(),
		blocks.NewBlockForTests(
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
				*pContentCommmit,
			),
			[]byte("this is some result"),
		),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(blockchain)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some additional bytes")
	retBlockchain, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytesAgain, err := adapter.ToBytes(retBlockchain)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !bytes.Equal(retBytes, retBytesAgain) {
		t.Errorf("the returned blockchain bytes are invalid")
		return
	}

}
