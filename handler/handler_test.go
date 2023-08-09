package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HomeHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Handler Return status %d expected 200", rr.Code)
	}
}

func TestSocketHandle(t *testing.T) {
	// req, err = http.NewRequest("Get")
}
