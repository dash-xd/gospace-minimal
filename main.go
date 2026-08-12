// Package function is the generic GCP Cloud Functions (2nd gen) HTTP
// entry point. It owns no routes or business logic, and never imports a
// specific chi-router-providing repo directly. It always imports the
// same placeholder module path below; go.mod's replace directive is
// what actually decides which repo backs it for a given deployment. See
// the replace line in go.mod to point this shell at a different router.
package function

import (
	"net/http"

	"gospace.invalid/router"
)

var r = router.NewRouter()

// Main is the exported entry point invoked by the Cloud Functions runtime.
func Main(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
