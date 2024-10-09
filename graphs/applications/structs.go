package applications

type schema struct {
	Head        head
	Points      []string
	Connections [][]byte // connection hash bytes
}

type head struct {
	Name    string
	Version uint
	Access  headAccess
}

type headAccess struct {
	Write accessWrite
	PRead *accessPermission
}

type accessWrite struct {
	Modify  accessPermission
	PReview *accessPermission
}

type accessPermission struct {
	Names        []string
	Compensation float64
}

type connection struct {
	Header connectionHeader
	Links  [][]byte // link hash bytes
	Suites [][]byte // suite hash bytes
}

type connectionHeader struct {
	Name     connectionName
	PReverse *connectionName
}

type connectionName struct {
	Name        string
	Cardinality connectionCardinality
}

type connectionCardinality struct {
	Min  uint
	PMax *uint
}

type link struct {
	Origin []byte // reference hash bytes
	Target []byte // reference hash bytes
}

type reference struct {
	Internal  string
	PExternal *external
}

type external struct {
	Schema string
	Point  string
}

type suite struct {
	Name         string
	Link         []byte // link hash bytes
	Expectations []suiteExpectation
}

type suiteExpectation struct {
	References [][]byte // reference hash bytes
	IsFail     bool
}
