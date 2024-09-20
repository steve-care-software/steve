package kinds

// Kind represents the point kind
type Kind interface {
	IsBytes() bool
	IsInt() bool
	IsUint() bool
	IsFloat() bool
	IsVector() bool
	Vector() Kind
}
