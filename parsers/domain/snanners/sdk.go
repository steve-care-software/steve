package scanners

// ElementFn represents the element func
type ElementFn func(values map[string]byte) (any, error)

// ElementsFn theub elements fn
type ElementsFn func(list []any) (any, error)

/*
// Scanner represents a scanner
type Scanner interface {
	Elements() Elements
}

// Elements represents elements
type Elements interface {
	List() []Element
}

// Element represents an element
type Element interface {
	Name() string
	Selector() []byte
	HasElements() bool
	Elements() Elements
}

// Selectors represents selectors
type Selectors interface {
	List() []Selector
}*/

// Elements represents the elements
type Elements interface {
	Name() string
	Selector() []byte
	ElementsFn() ElementsFn
}

// Element represents an element
type Element interface {
	Name() string
	Selector() []byte
	HasNext() bool
	Next()
}
