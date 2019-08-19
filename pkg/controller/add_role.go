package controller

import (
	"gitlab.com/tarantool/sandbox/tarantool-operator/pkg/controller/role"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, role.Add)
}