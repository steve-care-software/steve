package layers

type layers struct {
	list []Layer
}

func createLayers(
	list []Layer,
) Layers {
	out := layers{
		list: list,
	}

	return &out
}

// List returns the layers
func (obj *layers) List() []Layer {
	return obj.list
}
