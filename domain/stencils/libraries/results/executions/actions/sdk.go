package actions

// Builder represents the action builder
type Builder interface {
	Create() Builder
	IsSuccess() Builder
	IsPreviouslySaved() Builder
	Now() (Action, error)
}

// Action represents an execution action
type Action interface {
	IsSuccess() bool
	IsPreviouslySaved() bool
}
