package jsons

// Value represents a value
type Value struct {
	String string `json:"string"`
	Bytes  []byte `json:"bytes"`
}

// Content represents a content
type Content struct {
	String string   `json:"string"`
	Bytes  []byte   `json:"bytes"`
	Trees  []string `json:"trees"`
}

// LayerExecution represents a layer execution
type LayerExecution struct {
	IsStop     bool                      `json:"stop"`
	Assignment *LayerExecutionAssignment `json:"assignment"`
	Condition  *LayerExecutionCondition  `json:"condition"`
}

// LayerExecutionAssignment represents a layer execution assignment
type LayerExecutionAssignment struct {
	Name       string                   `json:"variable"`
	Assignable LayerExecutionAssignable `json:"assignable"`
}

// LayerExecutionAssignable represents a layer execution assignable
type LayerExecutionAssignable struct {
	Query   *LayerExecutionQuery  `json:"query"`
	Reduce  *LayerExecutionReduce `json:"reduce"`
	Compare []LayerConstantValue  `json:"compare"`
	Length  *LayerConstantValue   `json:"length"`
	Join    []LayerConstantValue  `json:"join"`
	Value   *LayerConstantValue   `json:"value"`
}

// LayerExecutionQuery represents a layer execution assignable query
type LayerExecutionQuery struct {
	Input  LayerConstantValue     `json:"input"`
	Layer  LayerInput             `json:"layer"`
	Values []LayerValueAssignment `json:"values"`
}

// LayerInput represents a layer input
type LayerInput struct {
	Variable string `json:"variable"`
	Layer    string `json:"layer"`
}

// LayerExecutionReduce represents a layer execution reduce
type LayerExecutionReduce struct {
	Variable string `json:"variable"`
	Length   uint8  `json:"length"`
}

// LayerConstantValue represents a layer constant value
type LayerConstantValue struct {
	Variable string `json:"variable"`
	Bytes    string `json:"bytes"`
}

// LayerExecutionCondition represents a layer execution condition
type LayerExecutionCondition struct {
	Variable   string           `json:"variable"`
	Executions []LayerExecution `json:"executions"`
}

// LayerReturn represents a layer return
type LayerReturn struct {
	Variable string          `json:"variable"`
	Kind     LayerReturnKind `json:"kind"`
}

// LayerReturnKind represents the layer return kind
type LayerReturnKind struct {
	IsPrompt   bool     `json:"isPrompt"`
	IsContinue bool     `json:"isContinue"`
	Execute    []string `json:"execute"`
}

// LayerValueAssignment represents the layer value assignment
type LayerValueAssignment struct {
	Name  string     `json:"name"`
	Value LayerValue `json:"value"`
}

// LayerValue represents the layer value
type LayerValue struct {
	Variable string `json:"variable"`
	Bytes    string `json:"bytes"`
	Layer    string `json:"layer"`
}

// LayerSuiteReturn represents the layer suite return
type LayerSuiteReturn struct {
	Output Content         `json:"output"`
	Kind   LayerReturnKind `json:"kind"`
}

// LayerSuite represents the layer suites
type LayerSuite struct {
	Name   string                 `json:"name"`
	Input  string                 `json:"input"`
	Return LayerSuiteReturn       `json:"return"`
	Values []LayerValueAssignment `json:"values"`
}

// Layer represents the layer instance
type Layer struct {
	Input      string            `json:"input"`
	Executions []LayerExecution  `json:"executions"`
	Return     LayerReturn       `json:"return"`
	Parameters map[string]string `json:"parameters"`
	Suites     []LayerSuite      `json:"suites"`
}

// Pointer represents a pointer
type Pointer struct {
	Path   []string      `json:"path"`
	Symbol PointerSymbol `json:"symbol"`
}

// PointerSymbol represents a pointer symbol
type PointerSymbol struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

// LinkSuiteOrigin represents a link suite origin
type LinkSuiteOrigin struct {
	Output Content `json:"output"`
	Query  string  `json:"query"`
}

// LinkSuite represents link suite
type LinkSuite struct {
	Origin  LinkSuiteOrigin `json:"origin"`
	Symbols []Pointer       `json:"symbols"`
	Input   Content         `json:"input"`
	//Expectation Query           `json:"expectation"`
}

// LinkOrigin represents a link origin
type LinkOrigin struct {
	Symbol    Pointer              `json:"symbol"`
	Direction *LinkOriginDirection `json:"direction"`
}

// LinkOriginDirection represents a link origin direction
type LinkOriginDirection struct {
	Next     *LinkOrigin `json:"next"`
	Previous *LinkOrigin `json:"previous"`
}

// LinkExecution represents a link execution
type LinkExecution struct {
	Layer  LayerInput             `json:"layer"`
	Values []LayerValueAssignment `json:"values"`
}

// LinkPreparation represents a link preparation
type LinkPreparation struct {
	IsStop    bool                      `json:"isStop"`
	Load      *Pointer                  `json:"load"`
	Exists    *Pointer                  `json:"exists"`
	Condition *LinkPreparationCondition `json:"condition"`
}

// LinkPreparationCondition represents a link preparation condition
type LinkPreparationCondition struct {
	Variable     string            `json:"variable"`
	Preparations []LinkPreparation `json:"preparations"`
}

// Link represens a link
type Link struct {
	Input        string            `json:"input"`
	Origins      []LinkOrigin      `json:"origins"`
	Execution    LinkExecution     `json:"execution"`
	Preparations []LinkPreparation `json:"preparations"`
	Suites       []LinkSuite       `json:"suites"`
}

// Symbol represents a symbol
type Symbol struct {
	String string `json:"string"`
	Bytes  []byte `json:"bytes"`
	Layer  *Layer `json:"layer"`
	Link   *Link  `json:"link"`
}

// Library represents a library
type Library struct {
	Path    []string          `json:"path"`
	Symbols map[string]Symbol `json:"symbols"`
}
