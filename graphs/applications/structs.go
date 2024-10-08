package applications

type schema struct {
}

type connection struct {
	header connectionHeader
	links  [][]byte // link hash bytes
	suites [][]byte // suite hash bytes
}

type connectionHeader struct {
	name     connectionName
	pReverse *connectionName
}

type connectionName struct {
	name        string
	cardinality connectionCardinality
}

type connectionCardinality struct {
	min  uint
	pMAx *uint
}

type link struct {
	origin []byte // reference hash bytes
	target []byte // reference hash bytes
}

type reference struct {
	internal  string
	pExternal *external
}

type external struct {
	schema string
	point  string
}

type suite struct {
	name         string
	link         []byte // link hash bytes
	expectations []suiteExpectation
}

type suiteExpectation struct {
	references [][]byte // reference hash bytes
	isFail     bool
}
