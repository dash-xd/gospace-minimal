// Not touched by the prep-router action - only source.go gets
// regenerated. This file owns wiring the router into an HTTP entry
// point so main.go doesn't have to.
package routersource

import "net/http"

// Serve builds this deployment's router once and returns the HTTP
// handler function the Cloud Functions runtime invokes per request.
func Serve() http.HandlerFunc {
	r := NewRouter()

	return func(w http.ResponseWriter, req *http.Request) {
		r.ServeHTTP(w, req)
	}
}
