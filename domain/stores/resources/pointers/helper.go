package pointers

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func uint64ToBytes(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, value)
	return bytes
}

func bytesToUint64(data []byte) (*uint64, error) {
	if len(data) != uint64Size {
		str := fmt.Sprintf("byte slice must be exactly %d bytes long to convert to uint64", uint64Size)
		return nil, errors.New(str)
	}

	value := binary.BigEndian.Uint64(data)
	return &value, nil
}
