package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Tests run from the package directory; change to project root
	// so that site/index.html and other assets are reachable.
	if err := os.Chdir("../.."); err != nil {
		panic("could not change to project root: " + err.Error())
	}
	os.Exit(m.Run())
}

func TestIndexReturns200(t *testing.T) {
	srv := httptest.NewServer(newHandler("site", "test-version"))
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
