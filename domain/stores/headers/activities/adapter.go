package activities

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits"
)

type adapter struct {
	commitsAdapter commits.Adapter
	hashAdapter    hash.Adapter
	builder        Builder
}

func createAdapter(
	commitsAdapter commits.Adapter,
	hashAdapter hash.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		commitsAdapter: commitsAdapter,
		hashAdapter:    hashAdapter,
		builder:        builder,
	}

	return &out
}

// ToBytes converts an activity to bytes
func (app *adapter) ToBytes(ins Activity) ([]byte, error) {
	retCommitsBytes, err := app.commitsAdapter.InstancesToBytes(ins.Commits())
	if err != nil {
		return nil, err
	}

	output := retCommitsBytes
	output = append(output, ins.Head().Bytes()...)
	return output, nil
}

// ToInstance converts bytes to activity
func (app *adapter) ToInstance(data []byte) (Activity, []byte, error) {
	retCommits, retRemaining, err := app.commitsAdapter.BytesToInstances(data)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) < hash.Size {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes to form the head hash, %d provided", hash.Size, len(retRemaining))
		return nil, nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(retRemaining)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.builder.Create().WithCommits(retCommits).WithHead(*pHash).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining[hash.Size:], nil
}
