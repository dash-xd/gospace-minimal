// Package function is the generic GCP Cloud Functions (2nd gen) HTTP
// entry point. It owns no routes or business logic of its own, and
// carries no knowledge of which chi router it's serving — that comes
// from newRouter, defined in source.go. See source.go for how a
// deployment picks its router.
package function

import "net/http"

var r = newRouter()

// Main is the exported entry point invoked by the Cloud Functions runtime.
func Main(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
