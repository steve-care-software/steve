package instructions

import (
	"errors"
	"fmt"
)

type elements struct {
	list []Element
}

func createElements(
	list []Element,
) Elements {
	out := elements{
		list: list,
	}

	return &out
}

// List returns the list of element
func (obj *elements) List() []Element {
	return obj.list
}

// Fetch fetches an element by index
func (obj *elements) Fetch(idx uint) (Element, error) {
	length := len(obj.list)
	if idx >= uint(length) {
		str := fmt.Sprintf("the provided index (%d) must be smaller than the length (%d) of the list", idx, length)
		return nil, errors.New(str)
	}

	return obj.list[idx], nil
}

// Value returns the value of the elements
func (obj *elements) Value() []byte {
	output := []byte{}
	for _, oneElement := range obj.list {
		output = append(output, oneElement.Value()...)
	}

	return output
}
