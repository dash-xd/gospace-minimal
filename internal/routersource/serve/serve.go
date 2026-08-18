// Package serve wires source.NewRouter() into an HTTP entry point so
// main.go doesn't have to. Unlike the source package, this file is never
// touched by the router action - only internal/routersource/source's
// single file gets dropped in.
package serve

import (
	"net/http"

	"github.com/dash-xd/gospace-minimal/internal/routersource/source"
)

// Serve builds this deployment's router once and returns the HTTP
// handler function the Cloud Functions runtime invokes per request.
func Serve() http.HandlerFunc {
	r := source.NewRouter()

	return func(w http.ResponseWriter, req *http.Request) {
		r.ServeHTTP(w, req)
	}
}
