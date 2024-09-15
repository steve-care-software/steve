package blockchains

import (
	"bytes"
	"math/big"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

func mine(
	hashAdapter hash.Adapter,
	transactions transactions.Transactions,
	requestedDifficulty uint8,
	miningValue uint8,
) ([]byte, error) {
	prefix := []byte{}
	casted := int(requestedDifficulty)
	for i := 0; i < casted; i++ {
		prefix = append(prefix, miningValue)
	}

	trxHash := transactions.Hash()
	cpt := big.NewInt(1)
	for {
		pHash, err := hashAdapter.FromMultiBytes([][]byte{
			trxHash.Bytes(),
			cpt.Bytes(),
		})

		if err != nil {
			return nil, err
		}

		if !bytes.HasPrefix(pHash.Bytes(), prefix) {
			cpt = cpt.Add(big.NewInt(1), cpt)
			continue
		}

		break
	}

	return cpt.Bytes(), nil
}
