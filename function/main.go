// Package function is the generic GCP Cloud Functions (2nd gen) HTTP
// entry point. It owns no routes or business logic, and never imports a
// chi-router-providing repo directly. It always imports the internal
// routersource wrapper below; that package's single file is what a
// deploy step regenerates to target a specific repo.
package function

import (
	newrouter "github.com/dash-xd/gospace-minimal/internal/routersource"
)

// Main is the exported entry point invoked by the Cloud Functions runtime.
var Main = newrouter.Serve()
