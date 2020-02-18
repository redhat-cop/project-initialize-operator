package controller

import (
	"github.com/redhat-cop/project-initialize-operator/project-initialize/pkg/controller/projectinitialize"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, projectinitialize.Add)
}
