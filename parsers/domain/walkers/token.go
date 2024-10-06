package walkers

type token struct {
	fn   ListFn
	next Walker
}

func createToken(
	fn ListFn,
) Token {
	return createTokenInternally(fn, nil)
}

func createTokenWithNext(
	fn ListFn,
	next Walker,
) Token {
	return createTokenInternally(fn, next)
}

func createTokenInternally(
	fn ListFn,
	next Walker,
) Token {
	out := token{
		fn:   fn,
		next: next,
	}

	return &out
}

// Fn returns the list fn
func (obj *token) Fn() ListFn {
	return obj.fn
}

// HasNext returns true if there is a next, false otherwise
func (obj *token) HasNext() bool {
	return obj.next != nil
}

// Next returns the next, if any
func (obj *token) Next() Walker {
	return obj.next
}
