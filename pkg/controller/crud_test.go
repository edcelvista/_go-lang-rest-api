package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCrudHandlerLIST(t *testing.T) {
	// Create a request to pass to the handler
	req := httptest.NewRequest(http.MethodGet, "/crud/list", nil)

	// Record the response
	rr := httptest.NewRecorder()

	CrudHandlerLIST(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expected := "Response"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())
	}
}
