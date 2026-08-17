package generate

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestGenerateDefault(t *testing.T) {
	src, err := Generate(Source{})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	if strings.Contains(string(src), "mount0 ") {
		t.Errorf("default output should not import any mounted package:\n%s", src)
	}

	if !strings.Contains(string(src), `mux.HandleFunc("GET /healthz"`) {
		t.Errorf("default output should still register the default route:\n%s", src)
	}

	mustParse(t, src)
}

func TestGenerateComposed(t *testing.T) {
	src, err := Generate(Source{Mounts: []Mount{
		{Package: "github.com/dash-xd/github-device-auth/router", Prefix: "/auth"},
		{Package: "github.com/dash-xd/some-other/router", Prefix: "/other"},
	}})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	if !strings.Contains(string(src), `mount0 "github.com/dash-xd/github-device-auth/router"`) {
		t.Errorf("composed output should import the first mounted package:\n%s", src)
	}

	if !strings.Contains(string(src), `mount1 "github.com/dash-xd/some-other/router"`) {
		t.Errorf("composed output should import the second mounted package:\n%s", src)
	}

	if !strings.Contains(string(src), `mux.Handle("/auth/", http.StripPrefix("/auth", mount0.NewRouter()))`) {
		t.Errorf("composed output should mount the first package under its prefix:\n%s", src)
	}

	if !strings.Contains(string(src), `mux.Handle("/other/", http.StripPrefix("/other", mount1.NewRouter()))`) {
		t.Errorf("composed output should mount the second package under its prefix:\n%s", src)
	}

	if !strings.Contains(string(src), `mux.HandleFunc("GET /healthz"`) {
		t.Errorf("composed output should still register the default route:\n%s", src)
	}

	mustParse(t, src)
}

func mustParse(t *testing.T, src []byte) {
	t.Helper()

	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "source.go", src, parser.AllErrors); err != nil {
		t.Fatalf("generated source.go does not parse: %v\n%s", err, src)
	}
}
