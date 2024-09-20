package suites

// Suites represents suites
type Suites interface {
	List() []Suite
}

// Suite represents a test suite
type Suite interface {
	Name() string
	Path() []string
	IsFail() bool
}
