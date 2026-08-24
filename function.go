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
	function "github.com/dash-xd/gospace-minimal/internal/function"
)

// Main is the exported entry point the Cloud Functions runtime invokes.
var Main = function.Main
