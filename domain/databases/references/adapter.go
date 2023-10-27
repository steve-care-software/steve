package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type adapter struct {
	contentKeysAdapter ContentKeysAdapter
	commitsAdapter     CommitsAdapter
	builder            Builder
}

func createAdapter(
	contentKeysAdapter ContentKeysAdapter,
	commitsAdapter CommitsAdapter,
	builder Builder,
) Adapter {
	out := adapter{
		contentKeysAdapter: contentKeysAdapter,
		commitsAdapter:     commitsAdapter,
		builder:            builder,
	}
	return &out
}

// ToContent converts reference to bytes
func (app *adapter) ToContent(ins Reference) ([]byte, error) {

	commitsBytes, err := app.commitsAdapter.ToContent(ins.Commits())
	if err != nil {
		return nil, err
	}

	commitLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(commitLengthBytes, uint64(len(commitsBytes)))

	output := []byte{}
	output = append(output, commitLengthBytes...)
	output = append(output, commitsBytes...)

	if ins.HasContentKeys() {
		contentKeyBytes, err := app.contentKeysAdapter.ToContent(ins.ContentKeys())
		if err != nil {
			return nil, err
		}

		contentKeysLengthBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(contentKeysLengthBytes, uint64(len(contentKeyBytes)))

		output = append(output, contentKeysLengthBytes...)
		output = append(output, contentKeyBytes...)
	}

	return output, nil
}

// ToReference converts bytes to reference
func (app *adapter) ToReference(content []byte) (Reference, error) {
	contentLength := len(content)
	if contentLength < minReferenceSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Reference instance, %d provided", minReferenceSize, contentLength)
		return nil, errors.New(str)
	}

	commitBytesLengthDelimiter := uint64(8)
	commitBytesLength := binary.LittleEndian.Uint64(content[:commitBytesLengthDelimiter])
	commitBytesDelimiter := int(commitBytesLength + commitBytesLengthDelimiter)
	if contentLength < commitBytesDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the Commits size of the Reference instance, %d provided", commitBytesDelimiter, contentLength)
		return nil, errors.New(str)
	}

	commits, err := app.commitsAdapter.ToCommits(content[commitBytesLengthDelimiter:commitBytesDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[commitBytesDelimiter:]
	builder := app.builder.Create().WithCommits(commits)
	if len(remaining) > 0 {
		contentKeysBytesLengthDelimiter := uint64(8)
		contentKeysBytesLength := binary.LittleEndian.Uint64(remaining[:contentKeysBytesLengthDelimiter])
		contentKeysBytesDelimiter := int(contentKeysBytesLength + contentKeysBytesLengthDelimiter)
		if contentLength < contentKeysBytesDelimiter {
			str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the ContentKeys size of the Reference instance, %d provided", contentKeysBytesDelimiter, contentLength)
			return nil, errors.New(str)
		}

		contentKeys, err := app.contentKeysAdapter.ToContentKeys(remaining[contentKeysBytesLengthDelimiter:contentKeysBytesDelimiter])
		if err != nil {
			return nil, err
		}

		builder.WithContentKeys(contentKeys)
	}

	return builder.Now()
}
