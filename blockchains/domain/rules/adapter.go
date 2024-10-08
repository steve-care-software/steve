package rules

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

type adapter struct {
	builder Builder
}

func createAdapter(
	builder Builder,
) Adapter {
	out := adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *adapter) ToBytes(ins Rules) ([]byte, error) {
	incrPerTrx := ins.IncreaseDifficultyPerTrx()
	incrPerTrxInBits := math.Float64bits(incrPerTrx)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, incrPerTrxInBits)

	output := []byte{
		ins.MiningValue(),
		ins.BaseDifficulty(),
	}

	output = append(output, bytes...)
	return output, nil
}

// ToInstance converts bytes to instance
func (app *adapter) ToInstance(data []byte) (Rules, []byte, error) {
	expected := floatSize + 2
	if len(data) < expected {
		str := fmt.Sprintf("the data was expected to be %d bytes, %d provided", expected, len(data))
		return nil, nil, errors.New(str)
	}

	bits := binary.LittleEndian.Uint64(data[2:expected])
	floatValue := math.Float64frombits(bits)

	retRules, err := app.builder.Create().
		WithMiningValue(data[0]).
		WithBaseDifficulty(data[1]).
		WithIncreaseDifficultyPerTrx(floatValue).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return retRules, data[expected:], nil
}
