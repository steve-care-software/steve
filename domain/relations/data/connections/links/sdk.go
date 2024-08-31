package links

// Links represents links
type Links interface {
	List() []Link
}

// Link represents a link
type Link interface {
	Name() string
	IsLeft() bool
}
