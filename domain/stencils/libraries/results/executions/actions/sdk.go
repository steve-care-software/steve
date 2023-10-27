package actions

// Action represents an execution action
type Action interface {
	IsSaved() bool
	IsPreviouslySaved() bool
}
