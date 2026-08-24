// Package function is the generic GCP Cloud Functions (2nd gen) HTTP
// entry point. It owns no routes or business logic, and never imports a
// chi-router-providing repo directly. It always imports the internal
// routersource/serve wrapper below, which in turn calls
// routersource/source - the one file a deploy step drops in to target a
// specific repo.
//
// It's not package main - Cloud Functions' Go buildpack generates its
// own main() and invokes Main by reflection.
//
// It lives at the module root (not under internal/, where it used to
// live) because that's where a deployed build actually needs it: the
// Go functions_framework buildpack (as of writing - see
// github.com/GoogleCloudPlatform/buildpacks/blob/main/cmd/go/functions_framework/lib/lib.go)
// discovers the target package by parsing whatever .go files sit
// directly at the root of the uploaded source, and never actually
// reads GOOGLE_FUNCTION_SOURCE despite it being documented as
// buildpack-agnostic - a subdirectory would be invisible to it.
package function

import (
	"github.com/dash-xd/gospace-minimal/internal/routersource/serve"
)

// Main is the exported entry point invoked by the Cloud Functions runtime.
var Main = serve.Serve()
