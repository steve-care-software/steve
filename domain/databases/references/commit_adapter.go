package references

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/steve/domain/hash"
)

type commitAdapter struct {
	hashAdapter   hash.Adapter
	actionAdapter ActionAdapter
	builder       CommitBuilder
}

func createCommitAdapter(
	hashAdapter hash.Adapter,
	actionAdapter ActionAdapter,
	builder CommitBuilder,
) CommitAdapter {
	out := commitAdapter{
		hashAdapter:   hashAdapter,
		actionAdapter: actionAdapter,
		builder:       builder,
	}

	return &out
}

// ToContent converts a Commit instance to content
func (app *commitAdapter) ToContent(ins Commit) ([]byte, error) {
	createdOnBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(createdOnBytes, uint64(ins.CreatedOn().UnixNano()))
	actionBytes, err := app.actionAdapter.ToContent(ins.Action())
	if err != nil {
		return nil, err
	}

	actionBytesAmount := make([]byte, 8)
	binary.LittleEndian.PutUint64(actionBytesAmount, uint64(len(actionBytes)))

	output := []byte{}
	output = append(output, createdOnBytes...)
	output = append(output, actionBytesAmount...)
	output = append(output, actionBytes...)
	if ins.HasParent() {
		parentHashBytes := ins.Parent().Bytes()

		output = append(output, 1)
		output = append(output, parentHashBytes...)
	}

	return output, nil
}

// ToCommit converts content to a Commit instance
func (app *commitAdapter) ToCommit(content []byte) (Commit, error) {
	contentLength := len(content)
	if contentLength < commitMinSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Commit instance, %d provided", commitMinSize, contentLength)
		return nil, errors.New(str)
	}

	createdOnDelimiter := 8
	createdOnUnixNano := binary.LittleEndian.Uint64(content[0:createdOnDelimiter])
	createdOn := time.Unix(0, int64(createdOnUnixNano)).UTC()

	actionBytesAmountDelimiter := createdOnDelimiter + 8
	actionBytesAmount := binary.LittleEndian.Uint64(content[createdOnDelimiter:actionBytesAmountDelimiter])

	actionBytesDelimiter := actionBytesAmountDelimiter + int(actionBytesAmount)
	action, err := app.actionAdapter.ToAction(content[actionBytesAmountDelimiter:actionBytesDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[actionBytesDelimiter:]
	builder := app.builder.Create().WithAction(action).CreatedOn(createdOn)
	if len(remaining) > 0 {
		pParentHash, err := app.hashAdapter.FromBytes(remaining[1:])
		if err != nil {
			return nil, err
		}

		builder.WithParent(*pParentHash)
	}

	return builder.Now()
}
