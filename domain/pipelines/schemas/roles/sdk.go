package roles

// Role represents role
type Role interface {
	HasRead() bool
	Read() []string
	HasWrite() bool
	Write() []string
	HasReview() bool
	Review() []string
}
