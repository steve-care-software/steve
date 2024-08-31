package points

type points struct {
	list []Point
}

func createPoints(
	list []Point,
) Points {
	out := points{
		list: list,
	}

	return &out
}

// List returns the list of points
func (obj *points) List() []Point {
	return obj.list
}
