package testinteg

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

var baseURL = flag.String("url", "http://localhost:8080", "base URL of the site to test")

func get(t *testing.T, path string) *http.Response {
	t.Helper()
	resp, err := http.Get(*baseURL + path)
	if err != nil {
		t.Fatalf("GET %s failed: %v", path, err)
	}
	return resp
}

func body(t *testing.T, resp *http.Response) string {
	t.Helper()
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}
	return string(b)
}

func TestSiteResponds(t *testing.T) {
	resp := get(t, "/")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestIndexContainsSiteName(t *testing.T) {
	resp := get(t, "/")
	if !strings.Contains(body(t, resp), "Global Autonomy") {
		t.Error("expected page to contain 'Global Autonomy'")
	}
}

func TestNoUnreplacedPlaceholders(t *testing.T) {
	resp := get(t, "/")
	if strings.Contains(body(t, resp), "__COMMIT_SHA__") {
		t.Error("__COMMIT_SHA__ placeholder was not replaced")
	}
}

func TestCSSServed(t *testing.T) {
	resp := get(t, "/css/styles.css")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 for CSS, got %d", resp.StatusCode)
	}
}

func TestOrganizationsPresent(t *testing.T) {
	resp := get(t, "/")
	b := body(t, resp)
	orgs := []string{"Intact Global", "Intaction", "Intact America"}
	for _, org := range orgs {
		if !strings.Contains(b, org) {
			t.Errorf("expected page to contain organization %q", org)
		}
	}
}

func TestJSServed(t *testing.T) {
	resp := get(t, "/js/nav.js")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 for nav.js, got %d", resp.StatusCode)
	}
}

func TestUsage(t *testing.T) {
	fmt.Printf("Running integration tests against: %s\n", *baseURL)
}
