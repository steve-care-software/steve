package lists

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToBytes converts list to bytes
func (app *adapter) ToBytes(list [][]byte) ([]byte, error) {
	lengthBytes := pointers.Uint64ToBytes(uint64(len(list)))
	output := lengthBytes
	for i := 0; i < len(list); i++ {
		amount := uint64(len(list[i]))
		amountBytes := pointers.Uint64ToBytes(amount)
		output = append(output, amountBytes...)
		output = append(output, list[i]...)
	}

	return output, nil
}

// ToInstance converts bytes to list
func (app *adapter) ToInstance(data []byte) ([][]byte, error) {
	if len(data) < pointers.Uint64Size {
		str := fmt.Sprintf(remainingTooSmallPatternErr, pointers.Uint64Size)
		return nil, errors.New(str)
	}

	pLength, err := pointers.BytesToUint64(data[:pointers.Uint64Size])
	if err != nil {
		return nil, err
	}

	list := [][]byte{}
	remaining := data[pointers.Uint64Size:]
	casted := int(*pLength)
	for i := 0; i < casted; i++ {
		if len(remaining) < pointers.Uint64Size {
			str := fmt.Sprintf(remainingTooSmallPatternErr, pointers.Uint64Size)
			return nil, errors.New(str)
		}

		pAmount, err := pointers.BytesToUint64(remaining[:pointers.Uint64Size])
		if err != nil {
			return nil, err
		}

		remaining = remaining[pointers.Uint64Size:]
		if uint64(len(remaining)) < *pAmount {
			str := fmt.Sprintf(remainingTooSmallPatternErr, pAmount)
			return nil, errors.New(str)
		}

		list = append(list, remaining[:*pAmount])
		remaining = remaining[*pAmount:]
	}

	return list, nil
}
