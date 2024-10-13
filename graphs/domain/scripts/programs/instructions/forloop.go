package instructions

type forLoop struct {
	index    ForIndex
	keyValue ForKeyValue
}

func createForLoopWithIndex(
	index ForIndex,
) ForLoop {
	return createFoorLoopInternally(index, nil)
}

func createForLoopWithKeyValue(
	keyValue ForKeyValue,
) ForLoop {
	return createFoorLoopInternally(nil, keyValue)
}

func createFoorLoopInternally(
	index ForIndex,
	keyValue ForKeyValue,
) ForLoop {
	out := forLoop{
		index:    index,
		keyValue: keyValue,
	}

	return &out
}

func (obj *forLoop) IsIndex() bool {
	return obj.index != nil
}

func (obj *forLoop) Index() ForIndex {
	return obj.index
}

func (obj *forLoop) IsKeyValue() bool {
	return obj.keyValue != nil
}

func (obj *forLoop) KeyValue() ForKeyValue {
	return obj.keyValue
}
