package references

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/trees"
)

type actionAdapter struct {
	hashAdapter     hash.Adapter
	hashTreeAdapter trees.Adapter
	builder         ActionBuilder
}

func createActionAdapter(
	hashAdapter hash.Adapter,
	hashTreeAdapter trees.Adapter,
	builder ActionBuilder,
) ActionAdapter {
	out := actionAdapter{
		hashAdapter:     hashAdapter,
		hashTreeAdapter: hashTreeAdapter,
		builder:         builder,
	}

	return &out
}

// ToContent converts an action to content
func (app *actionAdapter) ToContent(ins Action) ([]byte, error) {
	output := []byte{}
	if ins.HasInsert() {
		insert := ins.Insert()
		insertBytes, err := app.hashTreeAdapter.ToContent(insert)
		if err != nil {
			return nil, err
		}

		insertBytesAmount := make([]byte, 8)
		binary.LittleEndian.PutUint64(insertBytesAmount, uint64(len(insertBytes)))

		output = append(output, 0)
		output = append(output, insertBytesAmount...)
		output = append(output, insertBytes...)
	}

	if ins.HasDelete() {
		del := ins.Delete()
		delBytes, err := app.hashTreeAdapter.ToContent(del)
		if err != nil {
			return nil, err
		}

		output = append(output, 1)
		output = append(output, delBytes...)
	}

	return output, nil
}

// ToAction converts content to Action
func (app *actionAdapter) ToAction(content []byte) (Action, error) {
	contentLength := len(content)
	if contentLength < actionSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to an Action instance, %d provided", actionSize, contentLength)
		return nil, errors.New(str)
	}

	flag := content[0:1][0]
	remaining := content[1:]
	builder := app.builder.Create()
	if flag == 0 {
		insertBytesAmount := binary.LittleEndian.Uint64(remaining[:8])
		insertBytesDelimiter := 8 + insertBytesAmount
		htIns, err := app.hashTreeAdapter.ToHashTree(remaining[8:insertBytesDelimiter])
		if err != nil {
			return nil, err
		}

		builder.WithInsert(htIns)
		remaining = remaining[insertBytesDelimiter:]

		if len(remaining) > 0 {
			flag = remaining[0:1][0]
			remaining = remaining[1:]
		}
	}

	if flag == 1 {
		htIns, err := app.hashTreeAdapter.ToHashTree(remaining)
		if err != nil {
			return nil, err
		}

		builder.WithDelete(htIns)
	}

	return builder.Now()
}
