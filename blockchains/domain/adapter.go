package blockchains

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/blockchains/domain/blocks"
	"github.com/steve-care-software/steve/blockchains/domain/roots"
	"github.com/steve-care-software/steve/blockchains/domain/rules"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
	blockAdapter blocks.Adapter
	rulesAdapter rules.Adapter
	rootAdapter  roots.Adapter
	builder      Builder
}

func createAdapter(
	blockAdapter blocks.Adapter,
	rulesAdapter rules.Adapter,
	rootAdapter roots.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		blockAdapter: blockAdapter,
		rulesAdapter: rulesAdapter,
		rootAdapter:  rootAdapter,
		builder:      builder,
	}

	return &out
}

// ToBytes convert blockchain to bytes
func (app *adapter) ToBytes(ins Blockchain) ([]byte, error) {
	identifierBytes, err := ins.Identifier().MarshalBinary()
	if err != nil {
		return nil, err
	}

	name := ins.Name()
	nameLengthBytes := pointers.Uint64ToBytes(uint64(len(name)))

	description := ins.Description()
	descriptionLengthBytes := pointers.Uint64ToBytes(uint64(len(description)))

	rulesBytes, err := app.rulesAdapter.ToBytes(ins.Rules())
	if err != nil {
		return nil, err
	}

	rootBytes, err := app.rootAdapter.ToBytes(ins.Root())
	if err != nil {
		return nil, err
	}

	createdOnBytes := pointers.Uint64ToBytes(uint64(ins.CreatedOn().UTC().UnixNano()))

	output := identifierBytes
	output = append(output, nameLengthBytes...)
	output = append(output, []byte(name)...)
	output = append(output, descriptionLengthBytes...)
	output = append(output, []byte(description)...)
	output = append(output, rulesBytes...)
	output = append(output, rootBytes...)
	output = append(output, createdOnBytes...)
	if ins.HasHead() {
		headBytes, err := app.blockAdapter.InstanceToBytes(ins.Head())
		if err != nil {
			return nil, err
		}

		output = append(output, headBytes...)
	}

	return output, nil
}

// ToInstance convert bytes to blockchain
func (app *adapter) ToInstance(data []byte) (Blockchain, []byte, error) {
	if len(data) < uuidSize {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, uuidSize, len(data))
		return nil, nil, errors.New(str)
	}

	identidier, err := uuid.FromBytes(data[:uuidSize])
	if err != nil {
		return nil, nil, err
	}

	remaining := data[uuidSize:]
	if len(remaining) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(remaining))
		return nil, nil, errors.New(str)
	}

	nameBytesLength := data[uuidSize : uuidSize+pointers.Uint64Size]
	pNameLength, err := pointers.BytesToUint64(nameBytesLength)
	if err != nil {
		return nil, nil, err
	}

	remaining = remaining[pointers.Uint64Size:]
	if uint64(len(remaining)) < *pNameLength {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, *pNameLength, len(remaining))
		return nil, nil, errors.New(str)
	}

	name := remaining[:*pNameLength]
	remaining = remaining[*pNameLength:]

	if len(remaining) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(remaining))
		return nil, nil, errors.New(str)
	}

	pDescriptionLength, err := pointers.BytesToUint64(remaining[:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	remaining = remaining[pointers.Uint64Size:]
	if uint64(len(remaining)) < *pNameLength {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, *pNameLength, len(remaining))
		return nil, nil, errors.New(str)
	}

	description := remaining[:*pDescriptionLength]
	remaining = remaining[*pDescriptionLength:]

	retRules, retRemainingAfterRules, err := app.rulesAdapter.ToInstance(remaining)
	if err != nil {
		return nil, nil, err
	}

	retRoot, retRemainingAfterRoot, err := app.rootAdapter.ToInstance(retRemainingAfterRules)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemainingAfterRoot) < pointers.Uint64Size {
		str := fmt.Sprintf(dataLengthTooSmallErrPattern, pointers.Uint64Size, len(retRemainingAfterRoot))
		return nil, nil, errors.New(str)
	}

	pNanoTime, err := pointers.BytesToUint64(retRemainingAfterRoot[:pointers.Uint64Size])
	if err != nil {
		return nil, nil, err
	}

	creationTime := time.Unix(0, int64(*pNanoTime))
	retRemaining := retRemainingAfterRoot[pointers.Uint64Size:]
	builder := app.builder.Create().
		WithIdentifier(identidier).
		WithName(string(name)).
		WithDescription(string(description)).
		WithRules(retRules).
		WithRoot(retRoot).
		CreatedOn(creationTime)

	retHead, retRemainingAfterHead, err := app.blockAdapter.BytesToInstance(retRemaining)
	if err == nil {
		builder.WithHead(retHead)
		retRemaining = retRemainingAfterHead
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}
