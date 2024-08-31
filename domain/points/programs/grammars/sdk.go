package grammars

// Grammar represents the grammar
type Grammar interface {
	Name() string
	Description() string
	Code() []byte
}
