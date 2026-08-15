// Command genrouter regenerates internal/routersource/source.go, the one
// file main.go's import never varies from. Run it with -router-package to
// point this shell at a specific chi-router-providing package, or with no
// flags to reset it to the default stdlib mux.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dash-xd/gospace-minimal/internal/routersource/generate"
)

func main() {
	routerPackage := flag.String("router-package", "", "import path of a package exposing NewRouter() http.Handler; empty resets to the default stdlib mux")
	out := flag.String("out", "internal/routersource/source.go", "path to write the generated file to")
	flag.Parse()

	src, err := generate.Generate(generate.Source{RouterPackage: *routerPackage})
	if err != nil {
		fmt.Fprintln(os.Stderr, "genrouter:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*out, src, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "genrouter:", err)
		os.Exit(1)
	}
}
