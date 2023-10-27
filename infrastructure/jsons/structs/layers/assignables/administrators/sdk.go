package administrators

import "github.com/steve-care-software/steve/infrastructure/jsons/structs/values"

// Administrator represents administrator's assignables
type Administrator struct {
	Application *Application `json:"application"`
	Instance    *Instance    `json:"instance"`
}

// Application represents application
type Application struct {
	Authenticate *Authenticate `json:"authenticate"`
	Retrieve     []byte        `json:"retrieve"`
	Save         *Object       `json:"save"`
}

// Instance represents an instance
type Instance struct {
	Variable string         `json:"variable"`
	Method   InstanceMethod `json:"method"`
}

// Method represents the instance's method
type InstanceMethod struct {
	Administrator *AdministratorMethod `json:"administrator"`
	Identities    *IdentitiesMethod    `json:"identities"`
	Identity      *IdentityMethod      `json:"identity"`
}

// AdministratorMethod represents the administrator's method
type AdministratorMethod struct {
	HasIdentities bool `json:"hasIdentities"`
	GetIdentities bool `json:"getIdentities"`
}

// IdentitiesMethod represents identities method
type IdentitiesMethod struct {
	GetLength    bool `json:"getLength"`
	FetchAtIndex uint `json:"fetchAtIndex"`
}

// IdentityMethod represents identity method
type IdentityMethod struct {
	GetName      string `json:"getName"`
	GetContainer string `json:"getContainer"`
}

// Authenticate represents an authenticate
type Authenticate struct {
	Username values.Value `json:"username"`
	Password values.Value `json:"password"`
}

// Object represents an administrator's object
type Object struct {
	Identities []IdentityObject `json:"identities"`
}

// IdentityObject represents an identity object
type IdentityObject struct {
	Name      values.Value `json:"name"`
	Container values.Value `json:"container"`
}
