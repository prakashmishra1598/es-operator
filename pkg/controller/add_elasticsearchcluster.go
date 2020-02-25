package controller

import (
	"github.com/prakashmishra1598/es-operator/pkg/controller/elasticsearchcluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, elasticsearchcluster.Add)
}
