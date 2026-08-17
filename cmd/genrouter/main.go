// Command genrouter regenerates internal/routersource/source.go, the one
// file main.go's import never varies from. Run it with one or more -mount
// flags to compose router packages into this shell's default router, or
// with no flags to reset it to just the default route.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dash-xd/gospace-minimal/internal/routersource/generate"
)

func main() {
	var mounts mountFlags
	flag.Var(&mounts, "mount", "package=prefix pair to compose in, e.g. github.com/dash-xd/github-device-auth/router=/auth; repeatable")
	out := flag.String("out", "internal/routersource/source.go", "path to write the generated file to")
	flag.Parse()

	src, err := generate.Generate(generate.Source{Mounts: mounts})
	if err != nil {
		fmt.Fprintln(os.Stderr, "genrouter:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*out, src, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "genrouter:", err)
		os.Exit(1)
	}
}

// mountFlags collects repeated -mount flags into generate.Mount values.
type mountFlags []generate.Mount

func (m *mountFlags) String() string {
	if m == nil {
		return ""
	}

	parts := make([]string, len(*m))
	for i, mount := range *m {
		parts[i] = mount.Package + "=" + mount.Prefix
	}

	return strings.Join(parts, ",")
}

func (m *mountFlags) Set(value string) error {
	pkg, prefix, ok := strings.Cut(value, "=")
	if !ok || pkg == "" || prefix == "" {
		return fmt.Errorf("mount %q: want package=prefix", value)
	}

	*m = append(*m, generate.Mount{Package: pkg, Prefix: prefix})

	return nil
}
