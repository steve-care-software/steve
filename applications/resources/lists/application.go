package lists

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/applications/resources"
	"github.com/steve-care-software/steve/domain/stores/lists"
)

type application struct {
	resourceApp resources.Application
	listAdapter lists.Adapter
}

func createApplication(
	resourceApp resources.Application,
	listAdapter lists.Adapter,
) Application {
	out := application{
		resourceApp: resourceApp,
		listAdapter: listAdapter,
	}

	return &out
}

// Amount returns the amount of values contained in the list of that name
func (app *application) Amount(name string) (*uint, error) {
	values, err := app.RetrieveAll(name)
	if err != nil {
		return nil, err
	}

	length := uint(len(values))
	return &length, nil
}

// Retrieve retrieves a list of data
func (app *application) Retrieve(name string, index uint, amount uint) ([][]byte, error) {
	values, err := app.RetrieveAll(name)
	if err != nil {
		return nil, err
	}

	length := len(values)
	if length <= int(index) {
		str := fmt.Sprintf("the index (%d) is invalid because the list contain %d elements", index, length)
		return nil, errors.New(str)
	}

	toIndex := index + amount
	if length < int(toIndex) {
		str := fmt.Sprintf("the amount (%d) is invalid because the list contain %d elements, and the index + amount (%d) exceeds that length", amount, length, (index + amount))
		return nil, errors.New(str)
	}

	return values[index:toIndex], nil
}

// RetrieveAll returns all the data from the provided list
func (app *application) RetrieveAll(name string) ([][]byte, error) {
	data, err := app.resourceApp.Retrieve(name)
	if err != nil {
		return nil, err
	}

	return app.listAdapter.ToInstance(data)
}

// Append append values to the list
func (app *application) Append(name string, values [][]byte) error {
	list := values
	data, err := app.resourceApp.Retrieve(name)
	if err == nil {
		retList, err := app.listAdapter.ToInstance(data)
		if err != nil {
			return err
		}

		list = append(retList, values...)
	}

	updatedData, err := app.listAdapter.ToBytes(list)
	if err != nil {
		return err
	}

	return app.resourceApp.Save(name, updatedData)
}
