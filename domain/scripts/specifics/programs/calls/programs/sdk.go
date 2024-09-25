package programs

// Program represents a program call
type Program interface {
	Name() string
	Input() string
}
