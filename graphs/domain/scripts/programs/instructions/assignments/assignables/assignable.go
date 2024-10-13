package assignables

type assignable struct {
	engine      AssignableEngine
	listMap     ListMap
	programCall ProgramCall
	primitive   PrimitiveValue
	casting     Casting
	expand      Iterable
	operation   Operation
}

func createAssignableWithEngine(engine AssignableEngine) Assignable {
	return createAssignableWithInternally(engine, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithListMap(listMap ListMap) Assignable {
	return createAssignableWithInternally(nil, listMap, nil, nil, nil, nil, nil)
}

func createAssignableWithProgramCall(programCall ProgramCall) Assignable {
	return createAssignableWithInternally(nil, nil, programCall, nil, nil, nil, nil)
}

func createAssignableWithPrimitive(primitive PrimitiveValue) Assignable {
	return createAssignableWithInternally(nil, nil, nil, primitive, nil, nil, nil)
}

func createAssignableWithCasting(casting Casting) Assignable {
	return createAssignableWithInternally(nil, nil, nil, nil, casting, nil, nil)
}

func createAssignableWithExpand(expand Iterable) Assignable {
	return createAssignableWithInternally(nil, nil, nil, nil, nil, expand, nil)
}

func createAssignableWithOperation(operation Operation) Assignable {
	return createAssignableWithInternally(nil, nil, nil, nil, nil, nil, operation)
}

func createAssignableWithInternally(
	engine AssignableEngine,
	listMap ListMap,
	programCall ProgramCall,
	primitive PrimitiveValue,
	casting Casting,
	expand Iterable,
	operation Operation,
) Assignable {
	return &assignable{
		engine:      engine,
		listMap:     listMap,
		programCall: programCall,
		primitive:   primitive,
		casting:     casting,
		expand:      expand,
		operation:   operation,
	}
}

// IsEngine returns true if the assignable contains an engine
func (obj *assignable) IsEngine() bool {
	return obj.engine != nil
}

// Engine returns the engine if present
func (obj *assignable) Engine() AssignableEngine {
	return obj.engine
}

// IsListMap returns true if the assignable contains a list map
func (obj *assignable) IsListMap() bool {
	return obj.listMap != nil
}

// ListMap returns the list map if present
func (obj *assignable) ListMap() ListMap {
	return obj.listMap
}

// IsProgramCall returns true if the assignable contains a program call
func (obj *assignable) IsProgramCall() bool {
	return obj.programCall != nil
}

// ProgramCall returns the program call if present
func (obj *assignable) ProgramCall() ProgramCall {
	return obj.programCall
}

// IsPrimitive returns true if the assignable contains a primitive value
func (obj *assignable) IsPrimitive() bool {
	return obj.primitive != nil
}

// Primitive returns the primitive value if present
func (obj *assignable) Primitive() PrimitiveValue {
	return obj.primitive
}

// IsCasting returns true if the assignable contains a casting
func (obj *assignable) IsCasting() bool {
	return obj.casting != nil
}

// Casting returns the casting if present
func (obj *assignable) Casting() Casting {
	return obj.casting
}

// IsExpand returns true if the assignable contains an iterable
func (obj *assignable) IsExpand() bool {
	return obj.expand != nil
}

// Expand returns the iterable if present
func (obj *assignable) Expand() Iterable {
	return obj.expand
}

// IsOperation returns true if the assignable contains an operation
func (obj *assignable) IsOperation() bool {
	return obj.operation != nil
}

// Operation returns the operation if present
func (obj *assignable) Operation() Operation {
	return obj.operation
}
