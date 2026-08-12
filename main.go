// Package function is the generic GCP Cloud Functions (2nd gen) HTTP
// entry point. It owns no routes or business logic of its own: it just
// serves whatever chi router the imported package builds. To deploy a
// different chi router, swap the import below.
package function

import (
	"net/http"

	"github.com/dash-xd/github-device-auth/router"
)

var r = router.NewRouter()

// Main is the exported entry point invoked by the Cloud Functions runtime.
func Main(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
