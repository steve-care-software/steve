package lists

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToBytes converts list to bytes
func (app *adapter) ToBytes(list [][]byte) ([]byte, error) {
	output := []byte{}
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
	list := [][]byte{}
	remaining := data
	for {

		length := len(remaining)
		if length <= 0 {
			break
		}

		if len(remaining) < pointers.Uint64Size {
			str := fmt.Sprintf(remainijngTooSmallPatternErr, pointers.Uint64Size)
			return nil, errors.New(str)
		}

		pAmount, err := pointers.BytesToUint64(remaining[:pointers.Uint64Size])
		if err != nil {
			return nil, err
		}

		remaining = remaining[pointers.Uint64Size:]
		if uint64(len(remaining)) < *pAmount {
			str := fmt.Sprintf(remainijngTooSmallPatternErr, pAmount)
			return nil, errors.New(str)
		}

		list = append(list, remaining[:*pAmount])
		remaining = remaining[*pAmount:]
	}

	return list, nil
}