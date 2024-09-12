package uuids

import "github.com/google/uuid"

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// FromBytes converts bytes to ids
func (app *adapter) FromBytes(data []byte) ([]uuid.UUID, error) {
	list := []uuid.UUID{}
	amount := len(data) / UUIDSize
	for i := 0; i < amount; i++ {
		index := i * UUIDSize
		uuidBytes := data[index : index+UUIDSize]
		uuid, err := uuid.FromBytes(uuidBytes)
		if err != nil {
			return nil, err
		}

		list = append(list, uuid)
	}

	return list, nil
}

// FromInstances converts ids to bytes
func (app *adapter) FromInstances(list []uuid.UUID) ([]byte, error) {
	output := []byte{}
	for _, oneUuid := range list {
		uuidBytes, err := oneUuid.MarshalBinary()
		if err != nil {
			return nil, err
		}

		output = append(output, uuidBytes...)
	}

	return output, nil
}
