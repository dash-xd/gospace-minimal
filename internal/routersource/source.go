// Package routersource is a build-time drop-in: a deploy step regenerates
// this single file to point at whichever chi-router-providing repo should
// back this deployment. main.go's import of this package never changes;
// only the body below does. This is the current default, wired to
// github-device-auth's router.
package routersource

import (
	"net/http"

	router "github.com/dash-xd/github-device-auth/router"
)

// NewRouter returns the chi router for this deployment.
func NewRouter() http.Handler {
	return router.NewRouter()
}
