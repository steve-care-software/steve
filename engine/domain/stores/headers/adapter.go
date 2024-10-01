package headers

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities"
	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits/modifications/resources"
)

type adapter struct {
	resourcesAdapter  resources.Adapter
	activitiesAdapter activities.Adapter
	builder           Builder
}

func createAdapter(
	resourcesAdapter resources.Adapter,
	activitiesAdapter activities.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		resourcesAdapter:  resourcesAdapter,
		activitiesAdapter: activitiesAdapter,
		builder:           builder,
	}

	return &out
}

// ToBytes convert header to bytes
func (app *adapter) ToBytes(ins Header) ([]byte, error) {
	rootBytes, err := app.resourcesAdapter.InstancesToBytes(ins.Root())
	if err != nil {
		return nil, err
	}

	hasActivityBytes := byte(0)
	activityBytes := []byte{}
	if ins.HasActivity() {
		hasActivityBytes = byte(1)
		activityBytes, err = app.activitiesAdapter.ToBytes(ins.Activity())
		if err != nil {
			return nil, err
		}
	}

	output := rootBytes
	output = append(output, hasActivityBytes)
	output = append(output, activityBytes...)
	return output, nil
}

// ToInstance convert bytes to header
func (app *adapter) ToInstance(data []byte) (Header, []byte, error) {
	root, retRemaining, err := app.resourcesAdapter.BytesToInstances(data)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) < 1 {
		return nil, nil, errors.New("the data was expected to contain 1 byte that represents the flag that tells if there is an activity")
	}

	// true
	remaining := retRemaining[1:]
	builder := app.builder.Create().WithRoot(root)
	if retRemaining[0] == 1 {
		retActivity, retRemaining, err := app.activitiesAdapter.ToInstance(remaining)
		if err != nil {
			return nil, nil, err
		}

		remaining = retRemaining
		builder.WithActivity(retActivity)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}
