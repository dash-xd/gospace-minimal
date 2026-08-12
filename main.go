// Package function is the generic GCP Cloud Functions (2nd gen) HTTP
// entry point. It owns no routes or business logic of its own: it just
// serves whatever chi router the imported package builds. To deploy a
// different chi router, swap the import and the NewRouter call below.
package function

import (
	"net/http"

	deviceauthrouter "github.com/dash-xd/github-device-auth/router"
)

var router = deviceauthrouter.NewRouter()

// Main is the exported entry point invoked by the Cloud Functions runtime.
func Main(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
