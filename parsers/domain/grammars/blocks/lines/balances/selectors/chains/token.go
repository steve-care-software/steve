package chains

type token struct {
	index   uint
	element Element
}

func createToken(
	index uint,
) Token {
	return createTokenInternally(index, nil)
}

func createTokenWithElement(
	index uint,
	element Element,
) Token {
	return createTokenInternally(index, element)
}

func createTokenInternally(
	index uint,
	element Element,
) Token {
	out := token{
		index:   index,
		element: element,
	}

	return &out
}

// Index returns the index
func (obj *token) Index() uint {
	return obj.index
}

// HasElement returns true if there is an element, false otherwise
func (obj *token) HasElement() bool {
	return obj.element != nil
}

// Element returs the element, if any
func (obj *token) Element() Element {
	return obj.element
}
