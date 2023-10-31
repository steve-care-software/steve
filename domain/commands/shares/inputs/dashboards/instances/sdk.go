package instances

// Instance represents an instance command
type Instance interface {
	Object() string
	Content() Content
}

// Content reresents content
type Content interface {
	/*IsFetch() bool
	Fetch() fetches.Fetch
	IsInsert() bool
	Insert() inserts.Insert
	IsDelete() bool
	Delete() deletes.Delete*/
}
