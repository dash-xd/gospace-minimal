// Package function exists solely because the Go Cloud Functions
// buildpack discovers the deployed function's target package by
// parsing whatever .go files it finds directly at the root of the
// uploaded source - see internal/function's package doc for the full
// explanation, and why that package (this module's real implementation)
// can't be found there directly.
//
// Keep this file to a bare re-export: internal/function is the single
// source of truth for what Main actually does.
package function

import (
	"net/http"

	function "github.com/dash-xd/gospace-minimal/internal/function"
)

// Main is deliberately declared with the bare, unnamed function type
// here - not http.HandlerFunc, which is what internal/function.Main
// actually is and stays as (cmd/localserve needs exactly that for
// http.Server.Handler).
//
// The Go buildpack's generated main.go registers the function via a
// type assertion on an interface{} value:
//
//	if fnHTTP, ok := fn.(func(http.ResponseWriter, *http.Request)); ok {
//		funcframework.RegisterHTTPFunctionContext(ctx, "/", fnHTTP)
//	} else {
//		funcframework.RegisterEventFunctionContext(ctx, "/", fn)
//	}
//
// (github.com/GoogleCloudPlatform/buildpacks,
// cmd/go/functions_framework/lib/template_v1_1.go). A type assertion
// checks the exact dynamic type, unlike a plain assignment - boxing an
// http.HandlerFunc-typed value into that interface{} fails the
// assertion even though the underlying signature is identical, so the
// wrapper falls into the else branch and tries to register it as an
// event function instead. That registration then fails outright
// (RegisterEventFunctionContext requires a very different signature -
// see functions-framework-go/funcframework: "If fn has the wrong
// signature, RegisterEventFunction returns an error"), main() calls
// log.Fatalf, and the container exits before ever binding the PORT -
// which is exactly what surfaces as Cloud Run reporting the container
// "failed to start and listen on the port" within the health-check
// timeout, with nothing more specific in the logs.
//
// Declaring Main's static type here as the bare func type (Go allows
// this assignment - a named type converts to its compatible unnamed
// type in a var declaration, just not across an interface{} type
// assertion) makes the boxed dynamic type match what the assertion
// checks for, so the buildpack's wrapper takes the HTTP branch like it
// should.
var Main func(http.ResponseWriter, *http.Request) = function.Main
