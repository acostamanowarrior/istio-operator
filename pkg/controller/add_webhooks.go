package controller

import (
	"github.com/acostamanowarrior/istio-operator/pkg/controller/servicemesh/webhooks"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, webhooks.Add)
}
