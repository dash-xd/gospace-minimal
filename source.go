// This file is a deploy-time drop-in. main.go and the rest of this repo
// never change between deployments; this is the one file a deploy action
// replaces to point the generic entry point at a specific chi router.
//
// It only needs to exist, be package function, and implement:
//
//	func newRouter() http.Handler
//
// returning the target repo's router.NewRouter(). The version below is
// the current default, wired to github-device-auth.
package function

import (
	"net/http"

	"github.com/dash-xd/github-device-auth/router"
)

func newRouter() http.Handler {
	return router.NewRouter()
}
