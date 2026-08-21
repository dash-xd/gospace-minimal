package serve

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeUsesSourceNewRouterAsTheRouterFactory(t *testing.T) {
	h := Serve()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected the checked-in default source.NewRouter's /healthz to return 200, got %d", rec.Code)
	}
}
