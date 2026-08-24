// Package function is the generic GCP Cloud Functions (2nd gen) HTTP
// entry point. It owns no routes or business logic, and never imports a
// chi-router-providing repo directly. It always imports the internal
// routersource/serve wrapper below, which in turn calls
// routersource/source - the one file a deploy step drops in to target a
// specific repo.
//
// It's not package main - Cloud Functions' Go buildpack generates its
// own main() and invokes Main by reflection - so it lives under
// internal/ like the rest of this shell's implementation rather than
// cmd/. The deploy config points the buildpack at this directory via
// the GOOGLE_FUNCTION_SOURCE build environment variable, passed to
// `gcloud functions deploy --set-build-env-vars` (see
// .github/workflows/deploy-gcf.yml and, for the huram-abi-driven
// deploy of this specific router, terraform/device-auth-router in
// xd-dash/huram-abi) - not via a Terraform-managed GCS zip upload
// (google_cloudfunctions2_function's build_config.source only accepts
// a GCS object, and doesn't thread GOOGLE_FUNCTION_SOURCE through to
// the buildpack the same way `gcloud functions deploy --source=<local
// dir>` does).
package function

import (
	"github.com/dash-xd/gospace-minimal/internal/routersource/serve"
)

// Main is the exported entry point invoked by the Cloud Functions runtime.
var Main = serve.Serve()
