package externals

type external struct {
	schema string
	point  string
}

func createExternal(
	schema string,
	point string,
) External {
	out := external{
		schema: schema,
		point:  point,
	}

	return &out
}

// Schema returns the schema
func (obj *external) Schema() string {
	return obj.schema
}

// Point returns the point
func (obj *external) Point() string {
	return obj.point
}
