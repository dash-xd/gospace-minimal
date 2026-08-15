// Package routersource is a build-time drop-in: cmd/genrouter regenerates
// this single file to point at whichever chi-router-providing repo should
// back this deployment. main.go's import of this package never changes;
// only the body below does.
//
// Checked in, this is the default: a minimal stdlib http.ServeMux with a
// health check and nothing else, so this repo builds, tests, and deploys
// standalone with no external router dependency until something overrides
// it (see the .github/actions/router action, or run
// `go run ./cmd/genrouter -router-package=...` directly).
package routersource

import "net/http"

// NewRouter returns the router for this deployment.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return mux
}
