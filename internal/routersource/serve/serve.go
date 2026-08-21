// Package serve wires source.NewRouter() into an HTTP entry point so
// main.go doesn't have to. Unlike the source package, this file is never
// touched by the router action - only internal/routersource/source's
// single file gets dropped in.
package serve

import (
	"net/http"

	"github.com/dash-xd/gospace-minimal/internal/routersource/source"
)

// RouterFactory is the shape source.NewRouter must satisfy: build a
// complete http.Handler for this deployment, no arguments, called once.
// It's the same shape services that deploy through this shell build
// their own router with -- e.g. a service composing chi middleware with
// its own routes before handing back the finished http.Handler. This
// repo's drop-in is a file swapped in at build time (see
// .github/actions/router) rather than a function value passed at
// import time, but the contract it hands off to here -- a RouterFactory
// producing a plain http.Handler -- is the same either way, so a
// service's own router constructor needs no changes to be dropped in.
type RouterFactory func() http.Handler

// Serve builds this deployment's router once, via source.NewRouter as a
// RouterFactory, and returns the HTTP handler function the Cloud
// Functions runtime invokes per request. Assigning source.NewRouter to
// a RouterFactory-typed variable makes the compiler enforce that
// whatever's dropped into source.go still satisfies the contract.
func Serve() http.HandlerFunc {
	var factory RouterFactory = source.NewRouter
	r := factory()

	return func(w http.ResponseWriter, req *http.Request) {
		r.ServeHTTP(w, req)
	}
}
