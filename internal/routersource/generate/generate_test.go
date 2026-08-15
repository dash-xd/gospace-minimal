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

	if strings.Contains(string(src), "router \"") {
		t.Errorf("default output should not import a router package:\n%s", src)
	}

	mustParse(t, src)
}

func TestGenerateRemote(t *testing.T) {
	src, err := Generate(Source{RouterPackage: "github.com/dash-xd/github-device-auth/router"})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	if !strings.Contains(string(src), `router "github.com/dash-xd/github-device-auth/router"`) {
		t.Errorf("remote output should import the router package:\n%s", src)
	}

	if !strings.Contains(string(src), "return router.NewRouter()") {
		t.Errorf("remote output should delegate to router.NewRouter():\n%s", src)
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
