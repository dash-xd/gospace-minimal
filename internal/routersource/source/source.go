// Package source is a build-time drop-in: a deploy step overwrites this
// single file with whichever implementation should back this deployment,
// then builds the binary. It's the only file that ever changes - the
// serve package that wires it up lives elsewhere and is never touched by
// a deploy step. This file is plain Go - no codegen, no templating -
// so the dropped-in version is free to compose in as many router
// packages as it wants (mount several onto one mux, wrap them,
// whatever), all in ordinary Go code.
//
// Checked in, this is the default: a minimal stdlib http.ServeMux with a
// health check and nothing else, so this repo builds, tests, and deploys
// standalone with no external router dependency until something
// overwrites this file (see the .github/actions/router action).
package source

import "net/http"

// NewRouter returns the router for this deployment.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return mux
}
