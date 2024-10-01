package identities

import (
	"crypto/ed25519"
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources/pointers"
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

// ToBytes converts identity to bytes
func (app *adapter) ToBytes(ins Identity) ([]byte, error) {
	name := ins.Name()
	nameLengthBytes := pointers.Uint64ToBytes(uint64(len(name)))
	pkBytes := ins.PK().Seed()

	flags := []hash.Hash{}
	if ins.HasFlags() {
		flags = ins.Flags()
	}

	amountFlagBytes := pointers.Uint64ToBytes(uint64(len(flags)))
	output := nameLengthBytes
	output = append(output, []byte(name)...)
	output = append(output, pkBytes...)
	output = append(output, amountFlagBytes...)

	for _, oneFlag := range flags {
		output = append(output, oneFlag.Bytes()...)
	}

	return output, nil
}

// ToInstance converts bytes to identity
func (app *adapter) ToInstance(bytes []byte) (Identity, []byte, error) {
	if len(bytes) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(bytes))
		return nil, nil, errors.New(str)
	}

	pNameLength, err := pointers.BytesToUint64(bytes[:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	remaining := bytes[pointers.Uint64Size:]
	if uint64(len(remaining)) < *pNameLength {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, *pNameLength, len(remaining))
		return nil, nil, errors.New(str)
	}

	name := remaining[:*pNameLength]
	remaining = remaining[*pNameLength:]

	pk := ed25519.NewKeyFromSeed(remaining[:ed25519.SeedSize])
	remaining = remaining[ed25519.SeedSize:]
	if len(remaining) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(remaining))
		return nil, nil, errors.New(str)
	}

	pFlagAmount, err := pointers.BytesToUint64(remaining[:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	remaining = remaining[pointers.Uint64Size:]
	casted := int(*pFlagAmount)
	builder := app.builder.Create().WithName(string(name)).WithPK(pk)
	if casted > 0 {
		flags := []hash.Hash{}
		for i := 0; i < casted; i++ {
			if len(remaining) < hash.Size {
				str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(remaining))
				return nil, nil, errors.New(str)
			}

			pHash, err := app.hashAdapter.FromBytes(remaining[:hash.Size])
			if err != nil {
				return nil, nil, err
			}

			flags = append(flags, *pHash)
			remaining = remaining[hash.Size:]
		}

		builder.WithFlags(flags)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}
